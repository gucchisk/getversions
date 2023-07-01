/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	// "io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"golang.org/x/mod/semver"
	"github.com/spf13/cobra"
)

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
		fmt.Println("arg:" + args[0])
		resp, err := http.Get(args[0])
		if err != nil {
			fmt.Printf("error: %x\n", err)
		}
		defer resp.Body.Close()
		// b, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("error: %x\n", err)
		// }
		// fmt.Println(string(b))
		
		node, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Printf("error: %x\n", err)
		}
		//fmt.Printf("%x\n", node)
		nodes := findAll(node, atom.Img)

		for i := 1; i < len(nodes); i++ {
			n := nodes[i]
		// for _, n := range nodes {
			fmt.Printf("atom: %s\n", n.DataAtom.String())
			a := n.NextSibling.NextSibling
			if a != nil && a.DataAtom == atom.A {
				href, err := GetAttr(a, "href")
				if err != nil {
					fmt.Printf("%s", err)
				} else {
					version := strings.TrimRight(href, "/")
					if IsSemver(version) {
						fmt.Printf("version: %s\n", version)
					}
					
				}
			}
		}
	},
}

func IsSemver(version string) bool {
	v := version
	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}
	return semver.IsValid(v)
}

func GetAttr(node *html.Node, attr string) (string, error) {
	for _, a := range node.Attr {
		if a.Key == attr {
			return a.Val, nil
		}
	}
	return "", fmt.Errorf("%s is not found", attr)
}

func findFirst(node *html.Node, a atom.Atom) *html.Node {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.DataAtom == a {
			return c
		}
	}
	return nil
}

func findAll(node *html.Node, a atom.Atom) []*html.Node {
	var nodes []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			if c.DataAtom == a {
				nodes = append(nodes, c)
			}
			child_nodes := findAll(c, a)
			nodes = append(nodes, child_nodes...)
		}
	}
	return nodes
}

// func find(node *html.Node, a atom.Atom) {
// 	for c := node.FirstChild; c != nil; c = c.NextSibling {
// 		// fmt.Printf("%d\n", c.Type)
// 		if c.Type == html.ElementNode {
// 			fmt.Printf("node: %s parent: %s\n", c.DataAtom.String(), c.Parent.DataAtom.String())
// 			if c.DataAtom == a {
// 				fmt.Printf("%v\n", c.Data)
// 				if c.NextSibling.NextSibling != nil {
// 					fmt.Printf("next:%s\n", c.NextSibling.NextSibling.DataAtom.String())
// 				}
// 				// find(c, atom.A)
// 			}
// 			find(c, a)
// 		}
// 	}
// }

func init() {
	rootCmd.AddCommand(apacheCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
