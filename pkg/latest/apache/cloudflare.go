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

type Cloudflare struct {
	logger logr.Logger
}

func NewCloudflare() *Cloudflare {
	return &Cloudflare{}
}

func NewCloudflareWithLogger(logger logr.Logger) *Cloudflare {
	return &Cloudflare{logger: logger}
}

func (c *Cloudflare) GetLatestVersion(reader io.Reader, versionCondition string) (string, error) {
	node, err := html.Parse(reader)
	if err != nil {
		return "", err
	}
	nodes := htmlparser.FindAll(node, atom.A)
	c.logger.V(2).Info("", "len", len(nodes))
	latest := "v0.0.0"
	for i := 1; i < len(nodes); i++ {
		a := nodes[i]
		if a == nil || a.DataAtom != atom.A {
			continue
		}
		text := a.FirstChild.Data
		c.logger.V(2).Info("", "text", text)
		version := utils.ToSemver(strings.TrimRight(text, "/"))
		compareFunc := func(v string) {
			c.logger.V(1).Info("", "version", v)
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
