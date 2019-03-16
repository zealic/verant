package docker

import (
	"fmt"
	"regexp"
	"strings"

	version "github.com/hashicorp/go-version"
	"github.com/zealic/verant/internal"
)

const (
	baseURL          = "https://registry.hub.docker.com/v1/repositories"
	versionRegexpRaw = `^v?([0-9]+(\.[0-9]+)*?)$`
)

var (
	versionRegexp *regexp.Regexp
)

func init() {
	versionRegexp = regexp.MustCompile("^" + versionRegexpRaw + "$")
}

type ImageProvider struct {
	Owner     string
	Name      string
	Registry  string
	TagPrefix string
	TagSuffix string
}

type ImageTag struct {
	Name    string `json:"name"`
	Version string `json:"-"`
	Prefix  string `json:"-"`
	Suffix  string `json:"-"`
}

func (s *ImageTag) parse() {
	parts := strings.Split(s.Name, "-")
	switch len(parts) {
	case 1:
		if versionRegexp.MatchString(parts[0]) {
			s.Version = parts[0]
		}
	case 2:
		if versionRegexp.MatchString(parts[0]) {
			s.Version = parts[0]
			s.Suffix = parts[1]
		} else if versionRegexp.MatchString(parts[1]) {
			s.Prefix = parts[0]
			s.Version = parts[1]
		}
	default:
		// Match first version field
		for i, part := range parts {
			if versionRegexp.MatchString(part) {
				if i > 0 {
					s.Prefix = strings.Join(parts[:i], "-")
				}
				s.Version = part
				if len(parts) > i {
					s.Suffix = strings.Join(parts[i+1:], "-")
				}
				break
			}
		}
		// DO NOTHING
	}
}

func (s *ImageProvider) makeURL() string {
	imagePath := s.Name
	if s.Owner != "" {
		imagePath = s.Owner + "/" + imagePath
	}
	return fmt.Sprintf("%s/%s", baseURL, imagePath)
}

// GetTags all tags
func (s *ImageProvider) GetTags() ([]*ImageTag, error) {
	url := s.makeURL() + "/tags"
	tags := make([]*ImageTag, 0)
	err := internal.UnmarshalJSON(url, &tags)
	for _, tag := range tags {
		tag.parse()
	}
	return tags, err
}

func (s *ImageProvider) GetLastVersion() (*version.Version, error) {
	panic("NOT IMPL")
}
