package cwlogs_tailf

import (
	"flag"
	"log"
)

type CWLogsTailfParams struct {
	log_group_name  string
	log_stream_name string
}

func ParseFlag() (params *CWLogsTailfParams) {
	params = &CWLogsTailfParams{}

	flag.StringVar(&params.log_group_name, "g", "", "log group name")
	flag.StringVar(&params.log_stream_name, "s", "", "log stream name")
	flag.Parse()

	if params.log_group_name == "" {
		log.Fatal("'-g' is required")
	}

	if params.log_stream_name == "" {
		log.Fatal("'-s' is required")
	}

	return
}
