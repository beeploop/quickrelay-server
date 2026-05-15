package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
		Uptime int64  `json:"uptime_seconds"`
	}

	resp := Response{
		Status: "ok",
		Uptime: int64((time.Since(s.startTime)).Seconds()),
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
