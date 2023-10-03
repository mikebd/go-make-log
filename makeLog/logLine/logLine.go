package logLine

import (
	"encoding/json"
	"github.com/itchyny/timefmt-go"
	"github.com/xyproto/randomstring"
	"math/rand"
	"strings"
	"time"
)

const (
	maxCharactersPerWord = 8
	maxWordsPerMessage   = 15
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

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	randomstring.Seed()
}

func message() string {
	// TODO: Unsure why we sometimes get multiple consecutive spaces
	var sb strings.Builder
	sb.Grow((maxCharactersPerWord + 1) * maxWordsPerMessage)

	for i := 0; i < maxWordsPerMessage; i++ {
		if i > 0 {
			sb.WriteByte(' ')
		}
		randomWord := randomstring.HumanFriendlyEnglishString(random.Intn(maxCharactersPerWord))
		sb.WriteString(randomWord)
	}

	return sb.String()
}

func timestamp() string {
	const microsecondUtcFormat = "%Y-%m-%dT%H:%M:%S.%fZ"
	timeBuf := make([]byte, 0, 32)
	timeBuf = timefmt.AppendFormat(timeBuf, time.Now().UTC(), microsecondUtcFormat)
	return string(timeBuf)
}
