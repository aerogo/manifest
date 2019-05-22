package manifest

// Manifest represents a web manifest
type Manifest struct {
	Name            string         `json:"name"`
	ShortName       string         `json:"short_name"`
	Icons           []ManifestIcon `json:"icons,omitempty"`
	StartURL        string         `json:"start_url"`
	Display         string         `json:"display"`
	Lang            string         `json:"lang,omitempty"`
	ThemeColor      string         `json:"theme_color,omitempty"`
	BackgroundColor string         `json:"background_color,omitempty"`
}

// ManifestIcon represents a single icon in the web manifest.
type ManifestIcon struct {
	Source string `json:"src"`
	Sizes  string `json:"sizes"`
	Type   string `json:"type"`
}

// New creates a new manifest.
func New() *Manifest {
	return &Manifest{
		StartURL:  "/",
		Display:   "standalone",
		Lang:      "en",
		ShortName: "Untitled",
	}
}

// TODO: Warn about short name length (Google Lighthouse)
// https://developer.chrome.com/apps/manifest/name#short_name
