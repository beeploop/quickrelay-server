package api

import (
	"encoding/json"
	"net/http"

	"github.com/beeploop/quickrelay/internal/api/dto"
	"github.com/beeploop/quickrelay/internal/persistence"
	"github.com/google/uuid"
)

func (s *Server) handleSMS(w http.ResponseWriter, r *http.Request) {
	var payload dto.SendSmsRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := payload.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	job, err := persistence.InsertSmsJob(r.Context(), s.dbConn, persistence.NewSmsJob{
		ID:        uuid.NewString(),
		Recipient: payload.Recipient,
		Message:   payload.Message,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: notify dispatcher

	resp := dto.SendSmsResponse{
		JobID:  job.ID,
		Status: string(job.Status),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
