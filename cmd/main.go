package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/beeploop/quickrelay/internal/http"
)

func main() {
	PORT := ":3000"
	server := http.NewServer(PORT)

	go func() {
		log.Printf("quickrelay-server API listening on port: %s\n", PORT)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGTERM)
	<-quit

	log.Println("gracefully shutting down quickrelay-server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("could not gracefully shutdown server: %s\n", err.Error())
	}

	log.Println("quickrelay-server exited")
}
