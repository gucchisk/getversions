package apache

import (
	"io"
	"strings"

	"github.com/go-logr/logr"
	"github.com/gucchisk/getversions/pkg/utils/htmlparser"
	"github.com/gucchisk/getversions/utils"
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
	nodes := htmlparser.FindAll(node, atom.Img)
	a.logger.V(2).Info("", "len", len(nodes))
	latest := "v0.0.0"
	for i := 1; i < len(nodes); i++ {
		n := nodes[i]
		a.logger.V(2).Info("", "atom", n.DataAtom.String())
		a := n.NextSibling.NextSibling
		if a == nil || a.DataAtom != atom.A {
			continue
		}
		href, err := htmlparser.GetAttr(a, "href")
		if err != nil {
			// fmt.Printf("%s", err)
			continue
		}
		version := utils.ToSemver(strings.TrimRight(href, "/"))
		compareFunc := func(v string) {
			//logger.V(1).Info("", "version", v)
			if utils.IsBig(v, latest) {
				latest = v
			}
		}
		if versionCondition != "" {
			if strings.HasPrefix(version, utils.ToSemver(versionCondition)) {
				compareFunc(version)
			}
		} else {
			compareFunc(version)
		}
	}
	return latest, nil
}
