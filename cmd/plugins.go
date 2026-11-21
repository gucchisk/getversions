/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gucchisk/getversions/pkg/action"
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
    maven: Extract versions from Maven repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range actions {
			logger.V(2).Info("plugin found", "name", a.Name)
			action, err := getGetVersionAction(a)
			if err != nil {
				fmt.Printf("%s: error loading plugin - %s\n", a.Name, err.Error())
			} else {
				logger.V(2).Info("plugin", "action", action)
				pluginVersion, shortDescription, err := fetchPluginSummary(action)
				if err != nil {
					fmt.Printf("%s: error getting plugin info - %s\n", a.Name, err.Error())
					continue
				}
				fmt.Printf("%s@%s: %s\n", a.Name, pluginVersion, shortDescription)
			}
		}
	},
}

// fetchPluginSummary returns the version and short description used in the list.
func fetchPluginSummary(pluginAction *action.GetVersionsPluginActionRPC) (string, string, error) {
	pluginVersion, err := pluginAction.VersionWithError()
	if err != nil {
		return "", "", err
	}
	shortDescription, err := pluginAction.ShortWithError()
	return pluginVersion, shortDescription, err
}

func init() {
}
