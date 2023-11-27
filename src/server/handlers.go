package server

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func (s Server) RootHandler(w http.ResponseWriter, r *http.Request) {
	// print the list of files on the web page
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Files in ZIP-archive</h1>")
	fmt.Fprintln(w, "<ul>")
	for _, f := range s.Files {
		fmt.Fprintf(w, "<li>%s</li>\n", f)
	}
	fmt.Fprintln(w, "</ul>")
	// log the event
	s.Logger.Info("served file list", zap.String("remote_addr", r.RemoteAddr))
}
