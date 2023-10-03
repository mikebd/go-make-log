package makeLog

import (
	"github.com/xyproto/randomstring"
	"go-make-log/config"
	"go-make-log/makeLog/logLine"
	"log"
)

func MakeLog(args config.Arguments) error {
	randomstring.Seed()
	// TODO: Argument for minimum size of log file to create
	for i := 1; i < 500; i++ {
		log.Println(logLine.JsonLogLine())
	}
	return nil
}
