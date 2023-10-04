package logLine

import (
	"math/rand"
	"time"
)

var freqLevel = map[string]int{
	"debug": 500,
	"info":  200,
	"warn":  75,
	"error": 25,
	"fatal": 1,
}

var maxLevelLength = func() int {
	n := 0
	for k := range freqLevel {
		if len(k) > n {
			n = len(k)
		}
	}
	return n
}()

var randomLevel *rand.Rand

func init() {
	randomLevel = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

// freqsumLevel is a sum of all the frequencies in the freqLevel map
var freqsumLevel = func() int {
	n := 0
	for _, v := range freqLevel {
		n += v
	}
	return n
}()

// level will pick a level, weighted by the frequency table
func level() string {
	target := randomLevel.Intn(freqsumLevel)
	selected := "debug"
	n := 0
	for k, v := range freqLevel {
		n += v
		if n >= target {
			selected = k
			break
		}
	}
	return selected
}
