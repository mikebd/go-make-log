package makeLog

import (
	"go-make-log/config"
	"go-make-log/makeLog/logLine"
	"log"
	"strings"
)

func MakeLog(args *config.Arguments) error {
	// TODO: Argument for minimum size of log file to create
	// TODO: Use args.OutputFilename as output file

	var lineGenerator = func() func() string {
		if strings.EqualFold(args.Format, "json") {
			return logLine.JsonLogLine
		}
		return logLine.DelimitedLogLine
	}()

	for i := 1; i < 10; i++ {
		log.Println(lineGenerator())
	}
	return nil
}
