package apache

import (
	"io"

	"github.com/gucchisk/getversions/pkg/utils/htmlparser"
	"github.com/gucchisk/getversions/pkg/utils/version"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type ApacheAction struct {
}

func (a *ApacheAction) Short() string {
	return "from apache auto index page"
}

func (a *ApacheAction) Long() string {
	return "apache command extract the version from a list of downloadable versions displayed using Apache AutoIndex, like the one at https://archive.apache.org/dist/maven/maven-3/."
}

func (a *ApacheAction) GetVersions(reader io.Reader) []string {
	node, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}
	nodes := htmlparser.FindAll(node, atom.A)
	versions := []string{}
	for i := 1; i < len(nodes); i++ {
		anchor := nodes[i]
		if a == nil || anchor.DataAtom != atom.A {
			continue
		}
		textNodes := htmlparser.FindAllTexts(anchor)
		for _, textNode := range textNodes {
			v, _ := version.SearchVersion(textNode.Data)
			if v != "" {
				versions = append(versions, version.ToSemver(v))
			}
		}
	}
	return versions
}
