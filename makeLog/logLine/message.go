package logLine

import (
	"github.com/xyproto/randomstring"
	"math/rand"
	"strings"
	"time"
)

const (
	maxCharactersPerWord = 8
	maxWordsPerMessage   = 15
)

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
