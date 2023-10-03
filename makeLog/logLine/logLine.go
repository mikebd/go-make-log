package logLine

import (
	"encoding/json"
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
