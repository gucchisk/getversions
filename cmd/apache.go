/*
Copyright © 2023 gucchisk
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/gucchisk/getversions/pkg/latest/apache"
	"github.com/gucchisk/getversions/utils"
	"github.com/spf13/cobra"
)

// var log logr.Logger

// var htmltxt = strings.NewReader(`
// <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
// <html>
//  <head>
//   <title>Index of /maven/maven-3</title>
//  </head>
//  <body>
// <h1>Index of /maven/maven-3</h1>
// <pre><img src="/icons/blank.gif" alt="Icon "> <a href="?C=N;O=D">Name</a>                    <a href="?C=M;O=A">Last modified</a>      <a href="?C=S;O=A">Size</a>  <a href="?C=D;O=A">Description</a><hr><img src="/icons/back.gif" alt="[PARENTDIR]"> <a href="/maven/">Parent Directory</a>                             -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.0.5/">3.0.5/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.1.1/">3.1.1/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.2.5/">3.2.5/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.3.9/">3.3.9/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.5.4/">3.5.4/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.6.3/">3.6.3/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.8.8/">3.8.8/</a>                  2023-03-14 11:46    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.9.0/">3.9.0/</a>                  2023-02-06 08:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.9.1/">3.9.1/</a>                  2023-03-18 09:52    -
// <hr></pre>
// </body></html>
// `)

// apacheCmd represents the apache command
var apacheCmd = &cobra.Command{
	Use:   "apache",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.V(1).Info("", "arg", args[0])
		resp, err := http.Get(args[0])
		if err != nil {
			fmt.Printf("error: %x\n", err)
		}
		defer resp.Body.Close()
		iv, _ := cmd.Flags().GetString("version")
		logger.V(2).Info("", "Server", resp.Header.Get("Server"))
		getter := apache.NewApacheWithLogger(logger)
		latest, err := getter.GetLatestVersion(resp.Body, iv)
		if err != nil {
			fmt.Printf("error: %x\n", err)
		}
		fmt.Printf("%s", utils.FromSemver(latest))
	},
}

func init() {
	rootCmd.AddCommand(apacheCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	apacheCmd.Flags().StringP("version", "v", "", "version to get")
	// apacheCmd.Flags().Int("log", 0, "log level")
}
