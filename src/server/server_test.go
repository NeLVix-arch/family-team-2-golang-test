package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewServer(t *testing.T) {
	// create a logger
	logger := zap.NewNop()
	// create a slice of files
	files := []string{"file1.c", "file2.c", "file3.c"}
	// create a port
	port := "8080"
	// create a server
	server := NewServer(logger, files, port)
	// check if the server is not nil
	assert.NotNil(t, server, "server should not be nil")
	// check if the server has the same logger, files and port as the parameters
	assert.Equal(t, server.Logger, logger, "server should have the same logger as the parameter")
	assert.Equal(t, server.Files, files, "server should have the same files as the parameter")
	assert.Equal(t, server.Port, port, "server should have the same port as the parameter")
}

func TestRun(t *testing.T) {
	// create a logger
	logger := zap.NewNop()
	// create a slice of files
	files := []string{"file1.c", "file2.c", "file3.c"}
	// create a port
	port := "8080"
	// create a server
	server := NewServer(logger, files, port)
	// run the server in a goroutine
	go server.Run()
	// create a test request
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// create a test recorder
	rr := httptest.NewRecorder()
	// get the handler from the server
	handler := http.HandlerFunc(server.Run)
	// serve the request
	handler.ServeHTTP(rr, req)
	// check the status code
	assert.Equal(t, rr.Code, http.StatusOK, "handler should return OK status")
	// check the response body
	expected := "<h1>Files in ZIP-archive</h1>\n<ul>\n<li>file1.c</li>\n<li>file2.c</li>\n<li>file3.c</li>\n</ul>\n"
	assert.Equal(t, rr.Body.String(), expected, "handler should return the list of files")
}
