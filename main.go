package main

import (
	"context"
	_ "embed"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/babariviere/short/internal/api"
	"github.com/babariviere/short/internal/db"
	"github.com/babariviere/short/internal/oas"
	"github.com/jackc/pgx/v5"
)

//go:embed openapi.yaml
var openapiFile []byte

var rapidocHtml = `
<!doctype html> <!-- Important: must specify -->
<html>
  <head>
    <meta charset="utf-8"> <!-- Important: rapi-doc uses utf8 characters -->
    <script type="module" src="https://unpkg.com/rapidoc/dist/rapidoc-min.js"></script>
  </head>
  <body>
    <rapi-doc spec-url="/docs/openapi.yaml" render-style="view"> </rapi-doc>
  </body>
</html>
`

func main() {
	// This is rude but at least we can configure the database URL securely.
	psqlUrl := os.Getenv("DATABASE_URL")
	if psqlUrl == "" {
		log.Fatalln("DATABASE_URL is not configured.")
	}

	serverUrl := os.Getenv("SERVER_URL")
	if serverUrl == "" {
		serverUrl = "http://localhost:8080"
	}

	serverUrl = strings.TrimSuffix(serverUrl, "/")

	psql, err := pgx.Connect(context.Background(), psqlUrl)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	queries := db.New(psql)
	handler := api.NewHandler(serverUrl, queries)

	srv, err := oas.NewServer(handler)
	if err != nil {
		log.Fatalf("failed te create OpenAPI server: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", srv)
	mux.Handle("/docs", serveContent([]byte(rapidocHtml), "text/html"))
	mux.Handle("/docs/openapi.yaml", serveContent(openapiFile, "application/yaml"))
	log.Println("Listening on :8080")
	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func serveContent(content []byte, contentType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", contentType)
		w.Write(content)
	})
}
