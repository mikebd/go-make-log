package makeLog

import (
	"github.com/alecthomas/units"
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
	newline := []byte{'\n'}

	minSize, err := units.ParseStrictBytes(args.Size)
	if err != nil {
		return err
	}

	currentSize := int64(0)

	for currentSize < minSize {
		line := logGenerator()
		currentSize += int64(len(line) + len(newline))

		_, err := file.WriteString(line)
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
