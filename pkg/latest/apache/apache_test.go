package apache_test

// This test requires network connection

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gucchisk/getversions/pkg/latest/apache"
)

func TestGetLatestVersionForMaven(t *testing.T) {
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

func TestGetLatestVersionForNodejs(t *testing.T) {
	url := "https://nodejs.org/download/release/"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	writer := bytes.NewBuffer(nil)

	// 17 -> 17.9.1
	v, err := apache.NewApache().GetLatestVersion(io.TeeReader(resp.Body, writer), "17")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected := "v17.9.1"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}

	// 17.7 -> 17.7.2
	v, err = apache.NewApache().GetLatestVersion(strings.NewReader(writer.String()), "17.7")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected = "v17.7.2"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}
}

func TestGetLatestVersionForGradle(t *testing.T) {
	url := "https://services.gradle.org/distributions/"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	writer := bytes.NewBuffer(nil)

	// 8 -> 8.6
	v, err := apache.NewApache().GetLatestVersion(io.TeeReader(resp.Body, writer), "8")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected := "v8.6"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}

	// 8.2 -> 8.2.1
	v, err = apache.NewApache().GetLatestVersion(strings.NewReader(writer.String()), "8.2")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected = "v8.2.1"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}

	// 7 -> 7.6.3
	v, err = apache.NewApache().GetLatestVersion(strings.NewReader(writer.String()), "7")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected = "v7.6.3"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}

	// 8.5 -> 8.5
	v, err = apache.NewApache().GetLatestVersion(strings.NewReader(writer.String()), "8.5")
	if err != nil {
		t.Errorf("%s", err)
	}
	expected = "v8.5"
	if v != expected {
		t.Errorf("v = %s (expect: %s)", v, expected)
	}

}
