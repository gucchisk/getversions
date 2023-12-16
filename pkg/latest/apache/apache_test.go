package apache_test

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gucchisk/getversions/pkg/latest/apache"
)

func TestGetLatestVersionForApache(t *testing.T) {
	url := "https://archive.apache.org/dist/maven/maven-3/"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	writer := bytes.NewBuffer(nil)

	// 3 -> 3.9.6
	v, err := apache.NewApache().GetLatestVersion(io.TeeReader(resp.Body, writer), "3")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected := "v3.9.6"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}

	// 3.8 -> 3.8.8
	v, err = apache.NewApache().GetLatestVersion(strings.NewReader(writer.String()), "3.8")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected = "v3.8.8"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}
}

func TestGetLatestVersionForCloudflare(t *testing.T) {
	url := "https://nodejs.org/download/release/"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	writer := bytes.NewBuffer(nil)

	// 17 -> 17.9.1
	v, err := apache.NewCloudflare().GetLatestVersion(io.TeeReader(resp.Body, writer), "17")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected := "v17.9.1"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}

	// 17.7 -> 17.7.2
	v, err = apache.NewCloudflare().GetLatestVersion(strings.NewReader(writer.String()), "17.7")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected = "v17.7.2"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}
}
