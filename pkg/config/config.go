package config

// Ownbrew configuration
type Config struct {
	// Config version
	Version string `json:"version" yaml:"version"`
	// Path to the executable symlinks
	BinDir string `json:"binDir,omitempty" yaml:"binDir,omitempty"`
	// Path to your project taps
	TapDir string `json:"tapDir,omitempty" yaml:"tapDir,omitempty"`
	// Path for the downloaded sources
	TempDir string `json:"tempDir,omitempty" yaml:"tempDir,omitempty"`
	// Path to the versioned executables
	CellarDir string `json:"cellarDir,omitempty" yaml:"cellarDir,omitempty"`
	// List of packages that should be installed
	Packages []Package `json:"packages,omitempty" yaml:"packages,omitempty"`
}
