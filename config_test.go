package datadoglogsgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSourceWhenIsEmpty(t *testing.T) {
	cfg := NewDatadogConfiguration("xD", "", nil)
	subject := cfg.getSource()
	assert.Equal(t, GO, subject)
}

func TestGetSourceWhenHasValue(t *testing.T) {
	cfg := NewDatadogConfiguration("xD", "xD", nil)
	subject := cfg.getSource()
	assert.Equal(t, "xD", subject)
}

func TestGetTagsWhenIsEmpty(t *testing.T) {
	cfg := NewDatadogConfiguration("xD", "xD", nil)
	subject := cfg.getDDTags()
	assert.Equal(t, "", subject)
}

func TestGetTagsWhenHasValue(t *testing.T) {
	cfg := NewDatadogConfiguration("xD", "xD", []string{"xD", "xD2"})
	subject := cfg.getDDTags()
	assert.Equal(t, "xD,xD2", subject)
}
