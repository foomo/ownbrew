package config

// Ownbrew configuration
type Config struct {
	// Config version
	Version string `json:"version" yaml:"version"`
	// Path to the executable symlinks
	BinDir string `json:"binDir" yaml:"binDir"`
	// Path to your project taps
	TapDir string `json:"tapDir" yaml:"tapDir"`
	// Path for the downloaded sources
	TempDir string `json:"tempDir" yaml:"tempDir"`
	// Path to the versioned executables
	CellarDir string `json:"cellarDir" yaml:"cellarDir"`
	// List of packages that should be installed
	Packages []Package `json:"packages,omitempty" yaml:"packages,omitempty"`
}
