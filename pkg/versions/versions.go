package versions

import (
	"fmt"
	"strings"

	"github.com/gucchisk/getversions/pkg/utils/version"
)

type Version struct {
	origin string
	semver string
}

type Versions struct {
	versions   []Version
	condition  string
	onlyLatest bool
}

func NewVersions(versions []string) *Versions {
	vers := make([]Version, len(versions))
	for i, v := range versions {
		vers[i] = Version{
			origin: v,
			semver: version.ToSemver(v),
		}
	}
	return &Versions{
		versions:   vers,
		condition:  "",
		onlyLatest: false,
	}
}

func getOrigins(versions []Version) []string {
	origins := make([]string, len(versions))
	for i, ver := range versions {
		origins[i] = ver.origin
	}
	return origins
}

func getOrigin(versions []Version, semver string) (string, error) {
	for _, ver := range versions {
		if ver.semver == semver {
			return ver.origin, nil
		}
	}
	return "", fmt.Errorf("version not found: %s", semver)
}

func (v *Versions) Filter(condition string) *Versions {
	if condition == "" {
		return v
	}
	v.condition = version.ToSemver(condition)
	return v
}

func (v *Versions) OnlyLatest() *Versions {
	v.onlyLatest = true
	return v
}

func (v *Versions) Get() []string {
	if v.onlyLatest {
		latest := "v0.0.0"
		compareFunc := func(cv Version) {
			if version.IsBig(cv.semver, latest) {
				latest = cv.semver
			}
		}

		for _, ver := range v.versions {
			if v.condition != "" {
				if strings.HasPrefix(ver.semver, v.condition) {
					compareFunc(ver)
				}
			} else {
				compareFunc(ver)
			}
		}
		if latest != "v0.0.0" {
			orig, _ := getOrigin(v.versions, latest)
			return []string{orig}
		}
		return []string{}
	}

	if v.condition == "" {
		return getOrigins(v.versions)
	}

	filtered := []Version{}
	for _, ver := range v.versions {
		if strings.HasPrefix(ver.semver, v.condition) {
			filtered = append(filtered, ver)
		}
	}
	return getOrigins(filtered)
}
