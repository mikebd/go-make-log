package makeLog

import (
	"go-make-log/config"
	"go-make-log/makeLog/logLine"
	"log"
)

func MakeLog(args *config.Arguments) error {
	// TODO: Argument for minimum size of log file to create
	for i := 1; i < 10; i++ {
		log.Println(logLine.JsonLogLine())
	}
	return nil
}
