package main

import (
	"github.com/xyproto/randomstring"
	"log"
)

func makeLog(args arguments) error {
	randomstring.Seed()
	// TODO: Argument for minimum size of log file to create
	for i := 1; i < 500; i++ {
		log.Println(jsonLogLine())
	}
	return nil
}
