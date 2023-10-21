package utils

import (
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
