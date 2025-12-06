/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/gucchisk/getversions/pkg/action"
	act "github.com/gucchisk/getversions/pkg/action"
	"github.com/gucchisk/getversions/pkg/utils"
	vers "github.com/gucchisk/getversions/pkg/versions"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger logr.Logger
var pluginMap map[string]plugin.Plugin = make(map[string]plugin.Plugin)
var hcLogger hclog.Logger
var rootCmd *cobra.Command

func CreateRootCmd(level int) {
	// zapr
	var config zap.Config
	if level == 0 {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.Encoding = "json"
		config.Level = zap.NewAtomicLevelAt(zapcore.Level(-level))
	}
	z, _ := config.Build()
	logger = zapr.NewLogger(z)
	rootCmd = &cobra.Command{
		Use:   "getversions",
		Short: "root command for getversions",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.V(2).Info("root called")
			return cmd.Usage()
		},
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("plugin", "p", false, "List plugins")
	rootCmd.PersistentFlags().Int("log", 0, "log level")
	addCommands()

	hcLogger = hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		// hclog.Warn = 4
		Level: hclog.Level(4 - level),
	})
	addPluginCommands(hcLogger)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(level int) {
	err := rootCmd.Execute()

	defer func() {
		for _, a := range actions {
			a.Client.Kill()
		}
	}()

	if err != nil {
		fmt.Printf("error:%v\n", err.Error())
		os.Exit(1)
	}
}

func init() {
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "GETVERSIONS_PLUGIN",
	MagicCookieValue: "hello",
}

var actions []PluginAction

type PluginAction struct {
	Name   string
	Client *plugin.Client
}

func addCommands() {
	rootCmd.AddCommand(pluginsCmd)
	rootCmd.AddCommand(apacheCmd)
}

func addPluginCommands(hcLogger hclog.Logger) {
	plugins := utils.SearchPlugins()
	for _, p := range plugins {
		name := strings.Split(filepath.Base(p), "-")[1]
		pluginMap[name] = &act.GetVersionsPlugin{}
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: handshakeConfig,
			Plugins:         pluginMap,
			Cmd:             exec.Command(p),
			Logger:          hcLogger,
		})
		a := PluginAction{Name: name, Client: client}
		actions = append(actions, a)
		command, err := createPluginActionCmd(a)
		if err != nil {
			logger.Error(err, "create plugin command error", "plugin", name)
			continue
		}
		rootCmd.AddCommand(command)
	}
}

func createActionCmd(name string, a act.GetVersionsAction) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: a.Short(),
		Long:  a.Long(),
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := http.Get(args[0])
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			// buf := new(bytes.Buffer)
			// io.Copy(buf, resp.Body)
			// body := buf.Bytes()
			verList := a.GetVersions(resp.Body, *cmd.Flags())
			onlyLatest, err := cmd.Flags().GetBool("latest")
			if err != nil {
				return err
			}
			condition, err := cmd.Flags().GetString("version")
			if err != nil {
				return err
			}
			versions := vers.NewVersions(verList)
			versions = versions.Filter(condition)
			if onlyLatest {
				versions = versions.OnlyLatest()
			}
			filtered := versions.Get()
			for _, v := range filtered {
				fmt.Printf("%s\n", v)
			}
			return nil
		},
	}
	cmd.Flags().StringP("version", "v", "", "version condition to filter")
	cmd.Flags().BoolP("latest", "l", false, "only show the latest version")
	return cmd
}

func createPluginActionCmd(a PluginAction) (*cobra.Command, error) {
	action, err := getGetVersionAction(a)
	if err != nil {
		return nil, err
	}
	return createActionCmd(a.Name, action), nil
}

func getGetVersionAction(a PluginAction) (action.GetVersionsPluginAction, error) {
	rpcClient, err := a.Client.Client()
	if err != nil {
		logger.Error(err, "Client error")
		fmt.Printf("error: %x\n", err)
		return nil, err
	}
	raw, err := rpcClient.Dispense(a.Name)
	if err != nil {
		logger.Error(err, "Dispence error")
		fmt.Printf("error: %x\n", err)
		return nil, err
	}
	return raw.(action.GetVersionsPluginAction), nil
}
