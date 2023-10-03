package config

import (
	"flag"
)

type Arguments struct {
	LogTimestamps  bool
	OutputFilename string
}

func ParseArguments() Arguments {
	args := Arguments{}

	flag.BoolVar(&args.LogTimestamps, "lt", false, "Log timestamps (UTC)")
	flag.StringVar(&args.OutputFilename, "o", "", "Output filename")
	flag.Parse()

	return args
}
