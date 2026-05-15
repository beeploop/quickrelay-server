package api

import (
	"fmt"
	"io"
	"net/http"
)

func (s *Server) handleSMS(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// decode request
	// validate
	// save to DB
	// notify dispatcher
	// return response

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("raw body: %s\n", string(body))
}
