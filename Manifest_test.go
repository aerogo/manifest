package manifest_test

import (
	"testing"

	"github.com/aerogo/manifest"
)

func TestNew(t *testing.T) {
	m := manifest.New()

	if m.StartURL != "/" {
		t.Fatal("Invalid start URL")
	}

	if m.Display != "standalone" {
		t.Fatal("Invalid display value")
	}
}

func TestFromFile(t *testing.T) {
	m, err := manifest.FromFile("testdata/example.json")

	if err != nil {
		t.Fatal(err)
	}

	if m.Name != "Untitled Application" {
		t.Fatal("Invalid app name")
	}

	if len(m.Icons) != 4 {
		t.Fatal("Invalid JSON in the web manifest")
	}
}

func TestFromFileNonExisting(t *testing.T) {
	_, err := manifest.FromFile("testdata/nonexisting")

	if err == nil {
		t.Fatal("Should error")
	}
}

func TestFromFileInvalidJSON(t *testing.T) {
	_, err := manifest.FromFile("testdata/invalid.json")

	if err == nil {
		t.Fatal("Should error")
	}
}
