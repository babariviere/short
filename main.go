package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/babariviere/short/internal/api"
	"github.com/babariviere/short/internal/db"
	"github.com/babariviere/short/internal/oas"
	"github.com/jackc/pgx/v5"
)

func main() {
	// This is rude but at least we can configure the database URL securely.
	psqlUrl := os.Getenv("DATABASE_URL")
	if psqlUrl == "" {
		log.Fatalln("DATABASE_URL is not configured.")
	}

	psql, err := pgx.Connect(context.Background(), psqlUrl)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	queries := db.New(psql)
	handler := api.NewHandler(queries)

	srv, err := oas.NewServer(handler)
	if err != nil {
		log.Fatalf("failed te create OpenAPI server: %v", err)
	}

	log.Println("Listening on :8080")
	if err = http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
