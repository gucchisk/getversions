/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pluginsCmd represents the plugins command
var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "plugins command for getversions",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range actions {
			logger.V(2).Info("plugin found", "name", a.Name)
			action, err := getGetVersionAction(a)
			if err != nil {
				continue
			}
			logger.V(2).Info("plugin", "action", action)
			fmt.Printf("%s: %s\n", a.Name, action.Short())
		}
	},
}

func init() {
}
