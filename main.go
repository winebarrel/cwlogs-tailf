package main

import (
	"cwlogs_tailf"
	"log"
)

func init() {
	log.SetFlags(0)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	params := cwlogs_tailf.ParseFlag()
	err := cwlogs_tailf.Tailf(params)

	if err != nil {
		log.Fatal(err)
	}
}
