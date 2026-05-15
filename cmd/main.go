package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/beeploop/quickrelay/internal/api"
	"github.com/beeploop/quickrelay/internal/config"
	"github.com/beeploop/quickrelay/internal/persistence"
)

func main() {
	dbConn, err := persistence.OpenConnection()
	if err != nil {
		log.Fatalf("failed to open db connection: %s\n", err.Error())
	}

	if err := persistence.InitializeSchema(dbConn); err != nil {
		log.Fatalf("failed to initialize db schema: %s\n", err.Error())
	}

	server := api.New(config.Load(), dbConn)

	go func() {
		log.Printf("quickrelay-server API listening on port: %s\n", config.Load().PORT)
		if err := server.Start(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGTERM)
	<-quit

	log.Println("gracefully shutting down quickrelay-server...")

	if err := server.Stop(context.Background()); err != nil {
		log.Fatalf("could not gracefully shutdown server: %s\n", err.Error())
	}

	log.Println("quickrelay-server exited")
}
