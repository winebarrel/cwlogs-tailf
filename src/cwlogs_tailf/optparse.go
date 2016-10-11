package cwlogs_tailf

import (
	"flag"
	"log"
)

type CWLogsTailfParams struct {
	log_group_name         string
	log_stream_name        string
	log_stream_name_prefix string
	verbose                bool
}

func ParseFlag() (params *CWLogsTailfParams) {
	params = &CWLogsTailfParams{}

	flag.StringVar(&params.log_group_name, "g", "", "log group name")
	flag.StringVar(&params.log_stream_name_prefix, "p", "", "log stream name prefix")
	flag.StringVar(&params.log_stream_name, "s", "", "log stream name")
	flag.BoolVar(&params.verbose, "V", false, "verbose output")
	flag.Parse()

	if params.log_group_name == "" {
		log.Fatal("'-g' is required")
	}

	return
}
