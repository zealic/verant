package github

import (
	"fmt"

	"github.com/zealic/verant/internal"

	version "github.com/hashicorp/go-version"
)

const baseURL = "https://api.github.com"

type ReleaseInfo struct {
	Name       string   `json:"name"`
	TagName    string   `json:"tag_name"`
	Draft      bool     `json:"draft"`
	PreRelease bool     `json:"prerelease"`
	Assets     []*Asset `json:"assets"`
}

type Asset struct {
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
	Size        int    `json:"size"`
	URL         string `json:"browser_download_url"`
}

type ReleaseProvider struct {
	Owner string
	Name  string
}

func (s *ReleaseProvider) GetReleases() ([]*ReleaseInfo, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases", baseURL, s.Owner, s.Name)
	infos := make([]*ReleaseInfo, 0)
	err := internal.UnmarshalJSON(url, &infos)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (s *ReleaseProvider) GetLastRelease() (*ReleaseInfo, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases/latest", baseURL, s.Owner, s.Name)
	info := &ReleaseInfo{}
	err := internal.UnmarshalJSON(url, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (s *ReleaseProvider) GetLastVersion() (*version.Version, error) {
	rel, err := s.GetLastRelease()
	if err != nil {
		return nil, err
	}
	return version.NewVersion(rel.Name)
}
