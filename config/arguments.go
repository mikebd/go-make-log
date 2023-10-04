package config

import (
	"flag"
)

type Arguments struct {
	LogTimestamps  bool
	Format         string
	OutputFilename string
	Size           string
}

func ParseArguments() Arguments {
	args := Arguments{}

	flag.BoolVar(&args.LogTimestamps, "lt", false, "Log timestamps (UTC)")
	flag.StringVar(&args.Format, "f", "delimited", "Output format (delimited, json)")
	flag.StringVar(&args.OutputFilename, "o", "", "Output filename")
	flag.StringVar(&args.Size, "s", "1MiB", "Minimum output size (1KB, 1MB, 1GB, etc.)")
	flag.Parse()

	return args
}
