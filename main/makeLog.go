package main

import (
	"encoding/json"
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

func jsonLogLine() string {
	// TODO: Monotonic increasing timestamp
	// TODO: Randomize level
	// TODO: Randomize message
	result, _ := json.Marshal(map[string]interface{}{
		"timestamp": "2021-10-13T18:02:00.000Z",
		"level":     "info",
		"message":   "GET /api/v1/applyboard/programs/1234",
	})

	return string(result)
}
