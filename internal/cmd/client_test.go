package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/ecosystem/tablo/internal/cmd"
)

func TestNewClient(t *testing.T) {
	root := NewClient()
	require.NotNil(t, root)
	assert.NotEmpty(t, root.Use)
	assert.NotEmpty(t, root.Short)
	assert.NotEmpty(t, root.Long)
	assert.False(t, root.SilenceErrors)
	assert.True(t, root.SilenceUsage)
}
