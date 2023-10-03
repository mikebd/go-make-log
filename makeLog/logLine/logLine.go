package logLine

import (
	"encoding/json"
)

func JsonLogLine() string {
	result, _ := json.Marshal(map[string]interface{}{
		"timestamp": timestamp(),
		"level":     level(),
		"message":   message(),
	})

	return string(result)
}
