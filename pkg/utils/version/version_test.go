package version_test

import (
	"testing"

	"github.com/gucchisk/getversions/pkg/utils/version"
)

func TestIsSemver(t *testing.T) {
	versions := []string{
		"v0.0.0",
		"v0.0.1",
		"v1.0.0",
		"v9.99.999",
		"v1",
		"v1.1",
		"0",
		"0.0.1",
		"1.0.0",
	}
	for _, v := range versions {
		if !version.IsSemver(v) {
			t.Errorf("%s is not semver...\n", v)
		}
	}
}

func TestIsSemverInvalid(t *testing.T) {
	versions := []string{
		"v",
		"va",
		"test",
	}
	for _, v := range versions {
		if version.IsSemver(v) {
			t.Errorf("%s is semver...\n", v)
		}
	}
}

func TestFromSemver(t *testing.T) {
	tests := []struct {
		src string
		dst string
	}{
		{src: "v0.0.1", dst: "0.0.1"},
		{src: "v1", dst: "1"},
		{src: "1.0.0", dst: "1.0.0"},
		{src: "2", dst: "2"},
	}
	for _, test := range tests {
		v := version.FromSemver(test.src)
		if v != test.dst {
			t.Errorf("FromSemver(%s) = %s (expect: %s)\n", test.src, v, test.dst)
		}
	}
}
