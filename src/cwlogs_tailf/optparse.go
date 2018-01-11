package cwlogs_tailf

import (
	"flag"
	"github.com/tkuchiki/parsetime"
	"log"
	"regexp"
	"time"
)

type CWLogsTailfParams struct {
	log_group_name  string
	log_stream_name string
	filter          *regexp.Regexp
	start_time      int64
	verbose         bool
}

func parseTime(time_str string) int64 {
	p, err := parsetime.NewParseTime()

	if err != nil {
		log.Fatal(err)
	}

	t, err := p.Parse(time_str)

	if err != nil {
		log.Fatal(err)
	}

	return t.UnixNano() / int64(time.Millisecond)
}

func ParseFlag() (params *CWLogsTailfParams) {
	params = &CWLogsTailfParams{}
	var filter string
	var start_time string

	flag.StringVar(&params.log_group_name, "g", "", "log group name")
	flag.StringVar(&filter, "f", "", "log stream name filter regexp")
	flag.StringVar(&params.log_stream_name, "s", "", "log stream name")
	flag.StringVar(&start_time, "t", "", "start time")
	flag.BoolVar(&params.verbose, "V", false, "verbose output")
	flag.Parse()

	if params.log_group_name == "" {
		log.Fatal("'-g' is required")
	}

	if filter != "" {
		params.filter = regexp.MustCompile(filter)
	}

	if start_time != "" {
		params.start_time = parseTime(start_time)
	}

	return
}
