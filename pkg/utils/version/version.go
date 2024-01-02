package version

import (
	"regexp"
	"strings"

	"golang.org/x/mod/semver"
)

func IsSemver(version string) bool {
	v := ToSemver(version)
	return semver.IsValid(v)
}

func ToSemver(version string) string {
	v := version
	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}
	return v
}

func FromSemver(version string) string {
	v := version
	if strings.HasPrefix(v, "v") {
		v = v[1:len(v)]
	}
	return v
}

func IsBig(v1 string, v2 string) bool {
	return semver.Compare(v1, v2) == 1
}

func SearchVersion(str string) (string, error) {
	pattern := `\d+\.\d+`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	matches := re.FindAllString(str, -1)
	if len(matches) == 0 {
		return "", nil
	}
	patch, err := searchVersionPatch(str)
	if err != nil {
		return matches[0], err
	}
	if patch == "" {
		return matches[0], nil
	}
	return patch, nil
}

func searchVersionPatch(str string) (string, error) {
	pattern := `\d+\.\d+\.\d+`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	matches := re.FindAllString(str, -1)
	if len(matches) == 0 {
		return "", nil
	}
	return matches[0], nil
}
