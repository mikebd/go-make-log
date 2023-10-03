package main

import (
	"github.com/xyproto/randomstring"
	"log"
)

func makeLog(args arguments) error {
	randomstring.Seed()
	// TODO: Argument for minimum size of log file to create
	log.Println(jsonLogLine())
	log.Println(jsonLogLine())
	return nil
}
