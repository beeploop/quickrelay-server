package api

import (
	"context"
	"net/http"
	"time"

	"github.com/beeploop/quickrelay/internal/config"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	dbConn    *sqlx.DB
	mux       *http.ServeMux
	server    *http.Server
	startTime time.Time
}

func New(config *config.Config, db *sqlx.DB) *Server {
	mux := http.NewServeMux()

	s := &Server{
		dbConn: db,
		mux:    mux,
		server: &http.Server{
			Addr:    config.PORT,
			Handler: mux,
		},
		startTime: time.Now(),
	}

	s.registerRoutes()

	return s
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
