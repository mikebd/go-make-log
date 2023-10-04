package logLine

import (
	"github.com/itchyny/timefmt-go"
	"time"
)

var maxTimestampLength = func() int {
	println("maxTimestampLength")
	return len("2023-10-03T23:30:35.102988Z")
}()

func timestamp() string {
	const microsecondUtcFormat = "%Y-%m-%dT%H:%M:%S.%fZ"
	timeBuf := make([]byte, 0, 32)
	timeBuf = timefmt.AppendFormat(timeBuf, time.Now().UTC(), microsecondUtcFormat)
	return string(timeBuf)
}
