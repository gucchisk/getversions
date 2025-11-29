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
	tmpFlags := pflag.NewFlagSet("tmp", pflag.ContinueOnError)
	tmpFlags.ParseErrorsWhitelist.UnknownFlags = true
	tmpFlags.IntVar(&logLevel, "log", 0, "set log level (0=warn, 1=info, 2=debug)")
	tmpFlags.Parse(os.Args[1:])
	cmd.CreateRootCmd(logLevel)
	cmd.Execute(logLevel)
}
