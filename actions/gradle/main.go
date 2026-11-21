package main

import (
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/gucchisk/getversions/pluginapi"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/pflag"
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

func (g *GradleAction) FlagSet() pflag.FlagSet {
	return pflag.FlagSet{}
}

func (g *GradleAction) GetVersions(reader io.Reader, flags pflag.FlagSet) []string {
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

func main() {
	var pluginMap = map[string]plugin.Plugin{
		"gradle": &pluginapi.GetVersionsPlugin{Impl: &GradleAction{}},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginapi.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
