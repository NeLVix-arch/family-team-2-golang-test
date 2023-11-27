package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateLogger(t *testing.T) {
	// create a logger
	logger := CreateLogger()
	// check if the logger is not nil
	assert.NotNil(t, logger, "logger should not be nil")
	// check if the logger has the production level
	assert.Equal(t, logger.Core().Enabled(zap.DebugLevel), false, "logger should have production level")
}
