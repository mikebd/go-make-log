package makeLog

import (
	"go-make-log/config"
	"go-make-log/makeLog/logLine"
	"log"
	"os"
	"strings"
)

func MakeLog(args *config.Arguments) error {
	logGenerator := func() logLine.LogGenerator {
		if strings.EqualFold(args.Format, "json") {
			return logLine.JsonLogLine
		}
		return logLine.DelimitedLogLine
	}()

	file, err := os.OpenFile(args.OutputFilename, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	err = writeLogLines(args, file, logGenerator)
	if err != nil {
		return err
	}

	return nil
}

func writeLogLines(args *config.Arguments, file *os.File, logGenerator logLine.LogGenerator) error {
	// TODO: Argument for minimum size of log file to create

	newline := []byte{'\n'}

	for i := 1; i < 10; i++ {
		_, err := file.WriteString(logGenerator())
		if err != nil {
			return err
		}

		_, err = file.Write(newline)
		if err != nil {
			return err
		}
	}
	return nil
}
