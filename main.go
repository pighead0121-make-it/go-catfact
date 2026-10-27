package main

import (
	"context"
	"log"
	"os"
)

func main() {
	dbConn := os.Getenv("DATABASE_URL")
	listenAddr := os.Getenv("LISTEN_ADDR")

	ctx := context.Background()

	storage, err := NewPostgresStore(ctx, dbConn)
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.Init(ctx); err != nil {
		log.Fatal(err)
	}

	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewDBService(storage, svc)
	svc = NewLoggingService(svc)

	Server := NewAPIServer(svc)
	log.Printf("server starting on %s", listenAddr)
	log.Fatal(Server.Start(listenAddr))
}
