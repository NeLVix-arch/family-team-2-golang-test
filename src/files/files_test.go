package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestOpenZipFile(t *testing.T) {
	// create a logger
	logger := zap.NewNop()
	// create a file path
	file := "../../test.zip"
	// create an extension
	ext := ".c"
	// open the ZIP file
	files, close := OpenZipFile(logger, file, ext)
	// defer the closing function
	defer close()
	// check if the files are not nil
	assert.NotNil(t, files, "files should not be nil")
}
