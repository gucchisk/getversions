/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/gucchisk/getversions/cmd"
	"github.com/spf13/pflag"
)

func main() {
	var logLevel int
	if containsHelpFlag(os.Args[1:]) {
		exec(0)
		return
	}

	logFlags := pflag.NewFlagSet("log", pflag.ContinueOnError)
	logFlags.ParseErrorsWhitelist.UnknownFlags = true
	logFlags.IntVar(&logLevel, "log", 0, "set log level (0=warn, 1=info, 2=debug)")
	logFlags.Parse(os.Args[1:])
	exec(logLevel)
	// cmd.CreateRootCmd(logLevel)
	// cmd.Execute(logLevel)
}

func exec(logLevel int) {
	cmd.CreateRootCmd(logLevel)
	cmd.Execute(logLevel)
}

func containsHelpFlag(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}
