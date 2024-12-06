package config

import (
	"fmt"
	"strings"
)

type Package struct {
	// Name of the tap
	Tap string `json:"tap" yaml:"tap"`
	// Name of the package
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Names of the packages
	Names []string `json:"names,omitempty" yaml:"names,omitempty"`
	// Additional command args
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`
	// Version of the package to install
	Version string `json:"version" yaml:"version"`
	// List of tags
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func (c Package) AllNames() []string {
	names := c.Names
	if len(names) == 0 {
		names = append(names, c.Name)
	}
	return names
}

func (c Package) String() string {
	return fmt.Sprintf("Names: %s, Version: %s Tap: %s", c.AllNames(), c.Version, c.Tap)
}

func (c Package) URL() (string, error) {
	// foomo/tap/aws/kubectl
	parts := strings.Split(c.Tap, "/")
	if len(parts) < 4 {
		return "", fmt.Errorf("invalid tap format: %s", c.Tap)
	}
	return fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/ownbrew-%s/main/%s/%s.sh",
		parts[0],
		parts[1],
		parts[2],
		parts[3],
	), nil
}
