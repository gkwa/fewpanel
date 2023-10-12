package fewpanel

import (
	"flag"
)

var verbose bool
var logFormat string

func InitFlags() {
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output (shorthand)")
	flag.StringVar(&logFormat, "log-format", "", "Log format (text or json)")
}

func GetVerbose() bool {
	return verbose
}

func GetLogFormat() string {
	return logFormat
}
