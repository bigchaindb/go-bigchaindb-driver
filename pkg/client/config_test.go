package client

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoadFromReaders(t *testing.T) {
	cfgBytes, _ := ioutil.ReadFile("../fixtures/mock_config.json")

	var cfg ClientConfig
	err := Load(cfg, FromReaders("json", bytes.NewBuffer(cfgBytes)))

	assert.NoError(t, err, "Could not load config from readers")
}
