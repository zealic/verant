package github

import (
	"testing"

	version "github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetReleases(t *testing.T) {
	p := &ReleaseProvider{Owner: "hairyhenderson", Name: "gomplate"}
	vers, err := p.GetReleases()

	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(vers), 30)
}

func TestGetLastRelease(t *testing.T) {
	p := &ReleaseProvider{Owner: "hairyhenderson", Name: "gomplate"}
	rel, err := p.GetLastRelease()

	require.NoError(t, err)
	assert.Greater(t, len(rel.Assets), 0)
}

func TestGetLastVersion(t *testing.T) {
	p := &ReleaseProvider{Owner: "hairyhenderson", Name: "gomplate"}
	v, err := p.GetLastVersion()

	require.NoError(t, err)
	require.Condition(t, func() bool {
		return v.GreaterThan(version.Must(version.NewVersion("3.3.0")))
	})
}
