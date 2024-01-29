package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/babariviere/short/internal/db"
	"github.com/babariviere/short/internal/oas"
	"github.com/babariviere/short/internal/url"
	"github.com/jackc/pgx/v5"
)

var _ oas.Handler = (*handler)(nil)

type handler struct {
	serverUrl string
	queries   *db.Queries
}

func NewHandler(serverUrl string, queries *db.Queries) *handler {
	return &handler{
		serverUrl: serverUrl,
		queries:   queries,
	}
}

func (h *handler) RedirectLongURL(ctx context.Context, params oas.RedirectLongURLParams) (oas.RedirectLongURLRes, error) {
	res, err := h.queries.GetURLByHash(ctx, params.Hash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &oas.RedirectLongURLNotFound{}, nil
		}
		return nil, fmt.Errorf("unable to fetch URL by hash: %w", err)
	}

	return &oas.RedirectLongURLTemporaryRedirect{
		Location: oas.NewOptString(res.LongUrl),
	}, nil
}

func (h *handler) CreateShortURL(ctx context.Context, req *oas.CreateShortURLReq) (oas.CreateShortURLRes, error) {
	hash, err := url.GenerateHash(req.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to generate hash for url %s: %w", req.URL, err)
	}

	res, err := h.queries.InsertURL(ctx, db.InsertURLParams{
		Hash:    hash,
		LongUrl: req.URL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert url: %w", err)
	}

	return &oas.CreateShortURLCreated{
		Shorten: h.serverUrl + "/" + res.Hash,
	}, nil
}
