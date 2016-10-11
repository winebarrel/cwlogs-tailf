package cwlogs_tailf

import (
	"flag"
	"log"
	"regexp"
)

type CWLogsTailfParams struct {
	log_group_name  string
	log_stream_name string
	filter          *regexp.Regexp
	verbose         bool
}

func ParseFlag() (params *CWLogsTailfParams) {
	params = &CWLogsTailfParams{}
	var filter string

	flag.StringVar(&params.log_group_name, "g", "", "log group name")
	flag.StringVar(&filter, "f", "", "log stream name filter regexp")
	flag.StringVar(&params.log_stream_name, "s", "", "log stream name")
	flag.BoolVar(&params.verbose, "V", false, "verbose output")
	flag.Parse()

	if params.log_group_name == "" {
		log.Fatal("'-g' is required")
	}

	if filter != "" {
		params.filter = regexp.MustCompile(filter)
	}

	return
}
