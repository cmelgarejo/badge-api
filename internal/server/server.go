package server

import (
	"log"
	"net/http"
	"time"

	v1 "github.com/cmelgarejo/badge-api/internal/server/v1"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

// New inicialize a new server with configuration.
func New(port string) (*Server, error) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api/v1", v1.New())

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

// Close server resources.
func (serv *Server) Close() error {
	// TODO: add resource shutdown, if needed
	return nil
}

// Start the server.
func (serv *Server) Start() {
	log.Printf("Server running on %s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
