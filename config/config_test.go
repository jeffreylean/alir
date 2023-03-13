package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	config := new(Config)
	assert.Equal(t, int32(0), config.Port)

	path, err := filepath.Abs("../etc")
	assert.NoError(t, err)

	config.ConfigFile = path + "/sample.yaml"

	err = config.Load()
	assert.NoError(t, err)

	assert.Equal(t, int32(8080), config.Port)
}
