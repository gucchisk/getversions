package apache

import (
	"io"
	"strings"

	"github.com/go-logr/logr"
	"github.com/gucchisk/getversions/pkg/utils/htmlparser"
	"github.com/gucchisk/getversions/pkg/utils/version"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Apache struct {
	logger logr.Logger
}

func NewApache() *Apache {
	return &Apache{}
}

func NewApacheWithLogger(logger logr.Logger) *Apache {
	return &Apache{logger: logger}
}

func (a *Apache) GetLatestVersion(reader io.Reader, versionCondition string) (string, error) {
	node, err := html.Parse(reader)
	if err != nil {
		return "", err
	}
	nodes := htmlparser.FindAll(node, atom.A)
	a.logger.V(2).Info("", "len", len(nodes))
	latest := "v0.0.0"
	for i := 1; i < len(nodes); i++ {
		anchor := nodes[i]
		if a == nil || anchor.DataAtom != atom.A {
			continue
		}
		textNodes := htmlparser.FindAllTexts(anchor)
		a.logger.V(2).Info("", "textNodes", len(textNodes))
		var ver string
		for _, textNode := range textNodes {
			v, _ := version.SearchVersion(textNode.Data)
			a.logger.V(2).Info("", "text", textNode.Data)
			if v != "" {
				a.logger.V(2).Info("", "version", v)
				ver = version.ToSemver(v)
				break
			}
		}
		a.logger.V(2).Info("", "v", ver)
		compareFunc := func(v string) {
			a.logger.V(1).Info("", "version", v)
			if version.IsBig(v, latest) {
				latest = v
			}
		}
		if versionCondition != "" {
			if strings.HasPrefix(ver, version.ToSemver(versionCondition)) {
				compareFunc(ver)
			}
		} else {
			compareFunc(ver)
		}
	}
	return latest, nil
}
