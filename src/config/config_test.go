package config

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	// set the command line arguments
	os.Args = []string{"program", "-port", "8080", "-file", "./test.zip", "-ext", ".c"}
	// parse the config
	config := ParseConfig()
	// check if the config is not nil
	assert.NotNil(t, config, "config should not be nil")
	// check if the config has the correct values
	assert.Equal(t, config.Port, "8080", "config should have the correct port")
	assert.Equal(t, config.File, "./test.zip", "config should have the correct file")
	assert.Equal(t, config.Ext, ".c", "config should have the correct extension")
	// reset the flags
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
