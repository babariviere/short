package api_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/babariviere/short/internal/api"
	"github.com/babariviere/short/internal/db"
	"github.com/babariviere/short/internal/oas"
	"github.com/jackc/pgx/v5"
)

func prepare(t *testing.T) oas.Handler {
	name := fmt.Sprintf("test%d", time.Now().Unix())

	connCfg, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	root, err := pgx.ConnectConfig(ctx, connCfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		root.Close(ctx)
	})

	if _, err = root.Exec(ctx, "CREATE DATABASE "+name+";"); err != nil {
		t.Fatal("failed to create database", name)
	}
	t.Cleanup(func() {
		root.Exec(ctx, "DROP DATABASE "+name+";")
	})

	connCfg.Database = name
	conn, err := pgx.ConnectConfig(ctx, connCfg)
	if err != nil {
		log.Fatal(err)
	}
	t.Cleanup(func() {
		conn.Close(ctx)
	})

	schema, err := os.ReadFile("../../sql/schema.sql")
	if err != nil {
		t.Fatal(err)
	}
	conn.Exec(ctx, string(schema))

	return api.NewHandler("http://test", db.New(conn))
}

func TestCreateURL(t *testing.T) {
	h := prepare(t)

	cases := []string{
		"http://test.com",
		"http://github.com",
		"abcdef", // FIXME: should not pass but it works...
	}

	for _, url := range cases {
		shortenRes, err := h.CreateShortURL(context.TODO(), &oas.CreateShortURLReq{
			URL: url,
		})
		if err != nil {
			t.Errorf("failed to create url %q: %v", url, err)
			continue
		}

		shorten, ok := shortenRes.(*oas.CreateShortURLCreated)
		if !ok {
			t.Errorf("expected 201 status code, got 400 for url %q", url)
			continue
		}

		res, err := h.RedirectLongURL(context.TODO(), oas.RedirectLongURLParams{
			Hash: strings.TrimPrefix(shorten.Shorten, "http://test/"),
		})
		if err != nil {
			t.Errorf("failed to get redirect url for %q: %v", url, err)
			continue
		}

		redirect, ok := res.(*oas.RedirectLongURLTemporaryRedirect)
		if !ok {
			t.Errorf("expected 307 status code, got 404 for url %q with shorten %q", url, shorten.Shorten)
		}

		if redirect.Location.Value != url {
			t.Errorf("supplied %q but got %q during redirect", url, redirect.Location.Value)
		}
	}
}
