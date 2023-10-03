package main

import (
	"encoding/json"
	"github.com/itchyny/timefmt-go"
	"time"
)

func jsonLogLine() string {
	// TODO: Randomize level
	// TODO: Randomize message
	result, _ := json.Marshal(map[string]interface{}{
		"timestamp": timestamp(),
		"level":     "info",
		"message":   "GET /api/v1/applyboard/programs/1234",
	})

	return string(result)
}

func timestamp() string {
	const microsecondUtcFormat = "%Y-%m-%dT%H:%M:%S.%fZ"
	timeBuf := make([]byte, 0, 32)
	timeBuf = timefmt.AppendFormat(timeBuf, time.Now().UTC(), microsecondUtcFormat)
	return string(timeBuf)
}
