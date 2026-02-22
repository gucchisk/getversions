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
	Short: "list all installed plugins",
	Long: `List all installed getversions plugins and their descriptions.

  This command scans $GOPATH/bin for plugin binaries matching the pattern
  'getversions-*' and displays each plugin's name and short description.

  Plugins extend getversions functionality by adding support for additional
  version extraction sources. Each plugin registers as a subcommand that can
  be invoked directly.

  Example output:
    gradle: Extract versions from Gradle projects
    maven: Extract versions from Maven repositories

  These would replace the current placeholder text in cmd/plugins.go:15-21 to provide users with clear, helpful information when they run getversions plugins --help.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range actions {
			logger.V(2).Info("plugin found", "name", a.Name)
			action, err := getGetVersionAction(a)
			if err != nil {
				continue
			}
			logger.V(2).Info("plugin", "action", action)
			fmt.Printf("%s@%s: %s\n", a.Name, action.Version(), action.Short())
		}
	},
}

func init() {
}
