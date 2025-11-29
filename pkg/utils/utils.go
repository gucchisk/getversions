package utils

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"strings"
)

func GoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

func GoBin() string {
	return GoPath() + "/bin"
}

func SearchPlugins() []string {
	gobin := GoBin()
	files, _ := ioutil.ReadDir(gobin)
	plugins := []string{}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "getversions-") {
			plugins = append(plugins, fmt.Sprintf("%s/%s", gobin, f.Name()))
		}
	}
	return plugins
}
