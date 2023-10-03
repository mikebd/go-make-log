package logLine

import (
	"github.com/itchyny/timefmt-go"
	"time"
)

func timestamp() string {
	const microsecondUtcFormat = "%Y-%m-%dT%H:%M:%S.%fZ"
	timeBuf := make([]byte, 0, 32)
	timeBuf = timefmt.AppendFormat(timeBuf, time.Now().UTC(), microsecondUtcFormat)
	return string(timeBuf)
}
