package files

import (
	"archive/zip"
	"strings"

	"go.uber.org/zap"
)

// OpenZipFile opens a ZIP file and returns a slice of strings with the names of the files with the given extension and a closing function
func OpenZipFile(logger *zap.Logger, file string, ext string) ([]string, func()) {
	// open ZIP file
	zipReader, err := zip.OpenReader(file)
	if err != nil {
		logger.Fatal("failed to open ZIP file", zap.Error(err))
	}

	// get list of files with the given extension
	var files []string
	for _, f := range zipReader.File {
		if strings.HasSuffix(f.Name, ext) {
			files = append(files, f.Name)
		}
	}

	// return files and closing function
	return files, func() {
		zipReader.Close()
	}
}
