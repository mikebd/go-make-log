package logLine

import (
	"encoding/json"
	"github.com/itchyny/timefmt-go"
	"time"
)

func JsonLogLine() string {
	// TODO: Randomize level
	result, _ := json.Marshal(map[string]interface{}{
		"timestamp": timestamp(),
		"level":     "info",
		"message":   message(),
	})

	return string(result)
}

func timestamp() string {
	const microsecondUtcFormat = "%Y-%m-%dT%H:%M:%S.%fZ"
	timeBuf := make([]byte, 0, 32)
	timeBuf = timefmt.AppendFormat(timeBuf, time.Now().UTC(), microsecondUtcFormat)
	return string(timeBuf)
}
