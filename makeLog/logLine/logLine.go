package logLine

import (
	"encoding/json"
	"strings"
)

// LogGenerator is a function that returns a log line
type LogGenerator func() string

const delimiter = "|"

// DelimitedLogLine returns a log line in the format:
// 2023-10-03T23:30:35.102988Z|info |This is a log message
func DelimitedLogLine() string {
	level := level()
	message := message()
	sb := strings.Builder{}
	sb.Grow(maxTimestampLength + 1 + maxLevelLength + 1 + len(message))
	sb.WriteString(timestamp())
	sb.WriteString(delimiter)
	sb.WriteString(level)
	for i := len(level); i < maxLevelLength; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString(delimiter)
	sb.WriteString(message)
	return sb.String()
}

// JsonLogLine returns a log line in the format:
// {"level":"info","message":"This is a log message","timestamp":"2023-10-03T23:30:35.102988Z"}
func JsonLogLine() string {
	result, _ := json.Marshal(map[string]interface{}{
		"timestamp": timestamp(),
		"level":     level(),
		"message":   message(),
	})

	return string(result)
}
