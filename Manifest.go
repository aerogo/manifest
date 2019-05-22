package manifest

import (
	"os"

	jsoniter "github.com/json-iterator/go"
)

// Manifest represents a web manifest
type Manifest struct {
	Name            string         `json:"name"`
	ShortName       string         `json:"short_name"`
	Description     string         `json:"description"`
	Icons           []Icon         `json:"icons,omitempty"`
	StartURL        string         `json:"start_url"`
	Display         string         `json:"display"`
	Orientation     string         `json:"orientation,omitempty"`
	Language        string         `json:"lang,omitempty"`
	ThemeColor      string         `json:"theme_color,omitempty"`
	BackgroundColor string         `json:"background_color,omitempty"`
	TextDirection   string         `json:"dir,omitempty"`
	ServiceWorker   *ServiceWorker `json:"serviceworker,omitempty"`
	ScreenShots     []ScreenShot   `json:"screenshots,omitempty"`
}

// Icon represents a single icon in the web manifest.
type Icon struct {
	Source string `json:"src"`
	Sizes  string `json:"sizes"`
	Type   string `json:"type,omitempty"`
}

// ServiceWorker represents a service worker definition in the web manifest.
type ServiceWorker struct {
	Source         string `json:"src"`
	Scope          string `json:"scope"`
	UpdateViaCache string `json:"update_via_cache"`
}

// ScreenShot represents a screenshot definition in the web manifest.
type ScreenShot = Icon

// New creates a new manifest.
func New() *Manifest {
	return &Manifest{
		StartURL:  "/",
		Display:   "standalone",
		Language:  "en",
		ShortName: "Untitled",
	}
}

// FromFile creates a new manifest.
func FromFile(fileName string) (*Manifest, error) {
	// Open the file
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Decode JSON
	webManifest := New()
	decoder := jsoniter.NewDecoder(file)
	err = decoder.Decode(webManifest)

	if err != nil {
		return nil, err
	}

	return webManifest, nil
}

// TODO: Warn about short name length (Google Lighthouse)
// https://developer.chrome.com/apps/manifest/name#short_name
