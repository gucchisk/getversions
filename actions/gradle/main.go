package main

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/gucchisk/getversions/pkg/action"
	"github.com/hashicorp/go-plugin"
)

type GradleAction struct {
}

func (g *GradleAction) Version() string {
	return "1.0.0"
}

func (g *GradleAction) Short() string {
	return "Gradle release versions"
}

func (g *GradleAction) Long() string {
	return "Gradle release versions"
}

func (g *GradleAction) GetVersions(reader io.Reader) []string {
	fmt.Println("Fetching Gradle versions...")
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return []string{}
	}
	var versions = []string{}

	doc.Find("h3.u-text-with-icon > span + span").Each(func(i int, s *goquery.Selection) {
		versions = append(versions, s.Text())
	})

	return versions
}

var handShakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "GETVERSIONS_PLUGIN",
	MagicCookieValue: "hello",
}

func main() {
	// fmt.Println("Starting Gradle GetVersions plugin...")
	var pluginMap = map[string]plugin.Plugin{
		"gradle": &action.GetVersionsPlugin{Impl: &GradleAction{}},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handShakeConfig,
		Plugins:         pluginMap,
	})
}
