package cwlogs_tailf

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/cenkalti/backoff"
	"regexp"
	"sort"
	"time"
)

func backoffRateExceeded(operation func() error) (err error) {
	ticker := backoff.NewTicker(backoff.NewExponentialBackOff())

	for _ = range ticker.C {
		err = operation()

		if err == nil {
			ticker.Stop()
			break
		}

		awsErr, ok := err.(awserr.Error)

		if !ok {
			ticker.Stop()
			break
		}

		if awsErr.Code() != "ThrottlingException" {
			ticker.Stop()
			break
		}
	}

	return err
}

func formatTime(micro_sec int64) string {
	return time.Unix(0, micro_sec*int64(time.Millisecond)).Format(time.RFC3339)
}

func getLogEvents(svc *cloudwatchlogs.CloudWatchLogs, tailParams *CWLogsTailfParams) (err error) {
	params := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String(tailParams.log_group_name),
		LogStreamName: aws.String(tailParams.log_stream_name),
	}

	var next_token *string

	for {
		if next_token != nil {
			params.NextToken = next_token
		} else if tailParams.start_time > 0 {
			params.StartTime = aws.Int64(tailParams.start_time)
			params.StartFromHead = aws.Bool(true)
		}

		var resp *cloudwatchlogs.GetLogEventsOutput

		err = backoffRateExceeded(func() error {
			resp, err = svc.GetLogEvents(params)
			return err
		})

		if err != nil {
			break
		}

		for _, event := range resp.Events {
			if tailParams.verbose {
				fmt.Printf("%s\t%s\n", formatTime(*event.Timestamp), *event.Message)
			} else {
				fmt.Printf("%s\n", *event.Message)
			}
		}

		next_token = resp.NextForwardToken

		time.Sleep(1 * time.Second)
	}

	return
}

func geLogStreamNames(svc *cloudwatchlogs.CloudWatchLogs, log_group_name string, filter *regexp.Regexp, last_timestamp int64) (stream_names []string, err error) {
	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(log_group_name),
		Descending:   aws.Bool(true),
		OrderBy:      aws.String("LastEventTime"),
	}

	var next_token *string

	for {
		if next_token != nil {
			params.NextToken = next_token
		}

		var resp *cloudwatchlogs.DescribeLogStreamsOutput

		err = backoffRateExceeded(func() error {
			resp, err = svc.DescribeLogStreams(params)
			return err
		})

		if err != nil {
			return
		}

		for _, log_stream := range resp.LogStreams {
			if log_stream.LastIngestionTime == nil {
				continue
			}

			if *log_stream.LastIngestionTime < last_timestamp {
				return
			}

			if filter != nil && !filter.MatchString(*log_stream.LogStreamName) {
				continue
			}

			stream_names = append(stream_names, *log_stream.LogStreamName)
		}

		next_token = resp.NextToken

		if next_token == nil {
			break
		}
	}

	return
}

type NamedLogEvent struct {
	StreamName string
	Event      *cloudwatchlogs.OutputLogEvent
}

type NamedLogEvents []NamedLogEvent

func (events NamedLogEvents) Len() int {
	return len(events)
}

func (events NamedLogEvents) Swap(i, j int) {
	events[i], events[j] = events[j], events[i]
}

func (events NamedLogEvents) Less(i, j int) bool {
	return *events[i].Event.Timestamp < *events[j].Event.Timestamp
}

func getLogEventsFromLogGroup(svc *cloudwatchlogs.CloudWatchLogs, tailParams *CWLogsTailfParams) (err error) {
	var last_timestamp int64

	if tailParams.start_time > 0 {
		last_timestamp = tailParams.start_time
	} else {
		last_timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}

outer:
	for {
		var stream_names []string
		stream_names, err = geLogStreamNames(svc, tailParams.log_group_name, tailParams.filter, last_timestamp)

		if err != nil {
			break outer
		}

		var events NamedLogEvents

		for _, log_stream_name := range stream_names {
			params := &cloudwatchlogs.GetLogEventsInput{
				LogGroupName:  aws.String(tailParams.log_group_name),
				LogStreamName: aws.String(log_stream_name),
				StartTime:     aws.Int64(last_timestamp),
				StartFromHead: aws.Bool(true),
			}

			var resp *cloudwatchlogs.GetLogEventsOutput

			err = backoffRateExceeded(func() error {
				resp, err = svc.GetLogEvents(params)
				return err
			})

			if err != nil {
				break outer
			}

			for _, event := range resp.Events {
				if *event.Timestamp > last_timestamp {
					named_log_event := NamedLogEvent{log_stream_name, event}
					events = append(events, named_log_event)
				}
			}
		}

		sort.Sort(events)

		for _, named_log_event := range events {
			if tailParams.verbose {
				format_time := formatTime(*named_log_event.Event.Timestamp)
				fmt.Printf("%s\t%s\t%s\n", named_log_event.StreamName, format_time, *named_log_event.Event.Message)
			} else {
				fmt.Printf("%s\n", *named_log_event.Event.Message)
			}
			last_timestamp = *named_log_event.Event.Timestamp
		}

		time.Sleep(1 * time.Second)
	}

	return
}

func Tailf(params *CWLogsTailfParams) (err error) {
	svc := cloudwatchlogs.New(session.New())

	if params.log_stream_name != "" {
		err = getLogEvents(svc, params)
	} else {
		err = getLogEventsFromLogGroup(svc, params)
	}

	return
}
