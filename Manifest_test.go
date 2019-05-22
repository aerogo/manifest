package manifest_test

import (
	"testing"

	"github.com/aerogo/manifest"
)

func Test(t *testing.T) {
	m := manifest.New()

	if m.StartURL != "/" {
		t.Fatal("Invalid start URL")
	}

	if m.Display != "standalone" {
		t.Fatal("Invalid display value")
	}
}
