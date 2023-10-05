package logLine

import (
	"strings"
	"testing"
)

func Test_multipleConsecutiveSpaces(t *testing.T) {
	// Debug issue: Unsure why we sometimes get multiple consecutive spaces

	for i := 0; i < 100; i++ {
		message := message()
		if strings.Contains(message, "  ") {
			t.Fatalf("'%v' contains multiple spaces (iteration %d)", message, i)
		}
	}
}
