package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

// Server is a structure that holds the fields for the HTTP server
type Server struct {
	Logger *zap.Logger // logger for the server
	Files  []string    // slice of files to serve
	Port   string      // port for the server
}

// NewServer returns an instance of Server with the given parameters
func NewServer(logger *zap.Logger, files []string, port string) *Server {
	return &Server{
		Logger: logger,
		Files:  files,
		Port:   port,
	}
}

// Run creates and runs the HTTP server with the Chi router
func (s *Server) Run() {
	// create Chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Origin", "Content-Type"},
	}))

	// define handler for the root path
	r.Get("/", s.RootHandler)

	// start the HTTP server
	s.Logger.Info("starting HTTP server", zap.String("port", s.Port))
	if err := http.ListenAndServe(":"+s.Port, r); err != nil {
		s.Logger.Fatal("failed to start HTTP server", zap.Error(err))
	}
}
