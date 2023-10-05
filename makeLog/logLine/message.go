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

var randomMessage *rand.Rand

func init() {
	randomMessage = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	randomstring.Seed()
}

func message() string {
	var sb strings.Builder
	sb.Grow((maxCharactersPerWord + 1) * maxWordsPerMessage)

	wordGenerator := func() string {
		return randomstring.HumanFriendlyEnglishString(randomMessage.Intn(maxCharactersPerWord))
	}

	for i := 0; i < maxWordsPerMessage; i++ {
		if i > 0 {
			sb.WriteByte(' ')
		}
		randomWord := wordGenerator()
		for randomWord == "" {
			randomWord = wordGenerator()
		}
		sb.WriteString(randomWord)
	}

	return sb.String()
}
