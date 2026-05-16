package api

import "github.com/beeploop/quickrelay/internal/websocket"

func (s *Server) registerRoutes() {
	s.mux.HandleFunc("GET /", s.handleIndex)
	s.mux.HandleFunc("POST /sms", s.handleSMS)
	s.mux.HandleFunc("GET /ws", websocket.HandleConnection())
}
