package logLine

import (
	"encoding/json"
	"strconv"
	"strings"
)

// LogGenerator is a function that returns a log line
type LogGenerator func(lineNumber int64) string

const delimiter = "|"

// DelimitedLogLine returns a log line in the format:
// 2023-10-03T23:30:35.102988Z|info |This is a log message
func DelimitedLogLine(lineNumber int64) string {
	level := level()
	message := message()
	sb := strings.Builder{}
	sb.Grow(maxTimestampLength + 1 + maxLevelLength + 1 + len(message))
	sb.WriteString(timestamp())
	sb.WriteString(delimiter)
	if lineNumber > 0 {
		sb.WriteString(strconv.FormatInt(lineNumber, 10))
		sb.WriteString(delimiter)
	}
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
func JsonLogLine(lineNumber int64) string {
	jsonMap := map[string]interface{}{
		"timestamp": timestamp(),
		"level":     level(),
		"message":   message(),
	}
	if lineNumber > 0 {
		jsonMap["line"] = lineNumber
	}

	result, _ := json.Marshal(jsonMap)

	return string(result)
}
