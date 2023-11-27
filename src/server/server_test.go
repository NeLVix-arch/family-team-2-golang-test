package server_test

import (
	"family-team/src/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

// TestNewServer tests the NewServer function
func TestNewServer(t *testing.T) {
	// create a test logger
	logger := zaptest.NewLogger(t)

	// create a test slice of files
	files := []string{"file1.txt", "file2.jpg", "file3.pdf"}

	// create a test port
	port := "8080"

	// create a test server
	s := server.NewServer(logger, files, port)

	// assert that the server fields are correct
	assert.Equal(t, logger, s.Logger)
	assert.Equal(t, files, s.Files)
	assert.Equal(t, port, s.Port)
}

// TestRootHandler tests the RootHandler function
func TestRootHandler(t *testing.T) {
	// create a test logger
	logger := zaptest.NewLogger(t)

	// create a test slice of files
	files := []string{"file1.txt", "file2.jpg", "file3.pdf"}

	// create a test port
	port := "8080"

	// create a test server
	s := server.NewServer(logger, files, port)

	// create a test request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a test recorder
	rr := httptest.NewRecorder()

	// create a test router
	r := chi.NewRouter()
	r.Get("/", s.RootHandler)

	// serve the test request
	r.ServeHTTP(rr, req)

	// assert that the status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)
}
