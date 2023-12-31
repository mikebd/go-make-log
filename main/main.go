package main

import (
	"flag"
	"github.com/alecthomas/units"
	"go-make-log/config"
	"go-make-log/makeLog"
	"log"
	"os"
)

func main() {
	defer log.Println("Done")

	args := config.ParseArguments()
	initializeLogging(args.LogTimestamps)
	validateUsage(&args)

	err := run(&args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(args *config.Arguments) error {
	log.Println("Running:", os.Args)

	return makeLog.MakeLog(args)
}

func initializeLogging(logTimestamps bool) {
	if logTimestamps {
		log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(log.Lshortfile)
	}
	log.SetOutput(os.Stdout)
}

func validateUsage(args *config.Arguments) {
	var invalidUsage bool

	if len(args.OutputFilename) == 0 {
		invalidUsage = true
		log.Println("Missing output filename command line argument: -o <filename>")
	}

	_, err := units.ParseStrictBytes(args.Size)
	if err != nil {
		invalidUsage = true
		log.Println("Invalid size command line argument: -s <size>\n", err)
	}

	if invalidUsage {
		flag.Usage()
		os.Exit(2) // Same code used when a flag is provided but not defined
	}
}
