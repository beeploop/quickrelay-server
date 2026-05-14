package http

import "net/http"

func NewServer(port string) *http.Server {
	mux := http.NewServeMux()
	RegisterRoutes(mux)

	return &http.Server{
		Addr:    port,
		Handler: mux,
	}
}
