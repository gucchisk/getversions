package versions

import (
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

// matchesCondition はバージョンが条件に一致するかを、バージョンの区切り(".")単位で判定する。
// 例: 条件 "v1.2" は "v1.2" や "v1.2.3" に一致するが、"v1.20.0" には一致しない。
func matchesCondition(semver string, condition string) bool {
	return semver == condition || strings.HasPrefix(semver, condition+".")
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
		// v0.0.0 以下のバージョンは最新の候補にしない
		latest := Version{origin: "", semver: "v0.0.0"}
		found := false
		for _, candidate := range v.versions {
			if v.condition != "" && !matchesCondition(candidate.semver, v.condition) {
				continue
			}
			if version.IsBig(candidate.semver, latest.semver) {
				latest = candidate
				found = true
			}
		}
		if !found {
			return []string{}
		}
		return []string{latest.origin}
	}

	if v.condition == "" {
		return getOrigins(v.versions)
	}

	filtered := []Version{}
	for _, ver := range v.versions {
		if matchesCondition(ver.semver, v.condition) {
			filtered = append(filtered, ver)
		}
	}
	return getOrigins(filtered)
}
