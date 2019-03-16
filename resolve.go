package resolver

import (
	"strings"

	"github.com/zealic/verant/provider/docker"

	version "github.com/hashicorp/go-version"
)

// Provider Version provider
type Provider interface {
	GetLastVersion() (*version.Version, error)
}

func Resolve(spec *VersionSpec) Provider {
	name = strings.ToLower(name)
	switch name {
	case "alpine":
		return &docker.ImageProvider{Name: spec.Name}
	}

	return nil
}
