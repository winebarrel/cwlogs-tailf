package cwlogs_tailf

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"time"
)

func formatTime(micro_sec int64) string {
	sec := micro_sec / 1000
	nsec := micro_sec - sec
	return time.Unix(sec, nsec).Format(time.RFC3339)
}

func printMessages(events []*cloudwatchlogs.OutputLogEvent) {
	for _, event := range events {
		fmt.Printf("%s\t%s\n", formatTime(*event.Timestamp), *event.Message)
	}
}

func getLogEvents(svc *cloudwatchlogs.CloudWatchLogs, log_group_name string, log_stream_name string) (err error) {
	params := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String(log_group_name),
		LogStreamName: aws.String(log_stream_name),
	}

	var next_token *string

	for {
		if next_token != nil {
			params.NextToken = next_token
		}

		var resp *cloudwatchlogs.GetLogEventsOutput
		resp, err = svc.GetLogEvents(params)

		if err != nil {
			break
		}

		printMessages(resp.Events)
		next_token = resp.NextForwardToken

		time.Sleep(1 * time.Second)
	}

	return
}

func Tailf(params *CWLogsTailfParams) (err error) {
	svc := cloudwatchlogs.New(session.New())
	err = getLogEvents(svc, params.log_group_name, params.log_stream_name)
	return
}
