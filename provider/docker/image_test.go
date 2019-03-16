package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageTag_parse(t *testing.T) {
	assert := assert.New(t)
	var p *ImageTag
	set := func(tag string) {
		p = &ImageTag{Name: tag}
		p.parse()
	}

	set("9")
	assert.Equal("9", p.Version)

	set("9-slim")
	assert.Equal("9", p.Version)
	assert.Equal("slim", p.Suffix)

	set("debian-9-slim")
	assert.Equal("debian", p.Prefix)
	assert.Equal("9", p.Version)
	assert.Equal("slim", p.Suffix)

	set("1.2.3.4.5")
	assert.Equal("1.2.3.4.5", p.Version)

	set("apple-v9.1")
	assert.Equal("apple", p.Prefix)
	assert.Equal("v9.1", p.Version)

	set("super-mutant-v606.66.31-release-207")
	assert.Equal("super-mutant", p.Prefix)
	assert.Equal("v606.66.31", p.Version)
	assert.Equal("release-207", p.Suffix)
}

func TestGetTags(t *testing.T) {
	p := &ImageProvider{Name: "debian"}
	vers, err := p.GetTags()

	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(vers), 30)
}
