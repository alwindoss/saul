package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type (
	ServerOpt func(*Server)
)

func WithAddr(host, port string) ServerOpt {
	return func(s *Server) {
		s.host = host
		s.port = port
	}
}

func New(opts ...ServerOpt) *Server {
	const (
		defaultHost = "localhost"
		defaultPort = "8080"
	)
	s := &Server{
		host: defaultHost,
		port: defaultPort,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

type Server struct {
	host string
	port string
}

func (s *Server) Run() error {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Received at: %s", time.Now().String())
	})
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	log.Printf("Listening on %s", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		err = fmt.Errorf("exited unexpectedly from http.ListenAndServe: %w", err)
		return err
	}
	return nil
}
