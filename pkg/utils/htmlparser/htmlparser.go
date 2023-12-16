package htmlparser

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func GetAttr(node *html.Node, attr string) (string, error) {
	for _, a := range node.Attr {
		if a.Key == attr {
			return a.Val, nil
		}
	}
	return "", fmt.Errorf("%s is not found", attr)
}

func FindFirst(node *html.Node, a atom.Atom) *html.Node {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.DataAtom == a {
			return c
		}
	}
	return nil
}

func FindAll(node *html.Node, a atom.Atom) []*html.Node {
	var nodes []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		// logger.V(2).Info("", "type", c.Type)
		if c.Type == html.ElementNode {
			if c.DataAtom == a {
				nodes = append(nodes, c)
			}
			child_nodes := FindAll(c, a)
			nodes = append(nodes, child_nodes...)
		}
	}
	return nodes
}
