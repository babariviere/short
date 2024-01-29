package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/babariviere/short/internal/db"
	"github.com/babariviere/short/internal/oas"
	"github.com/jackc/pgx/v5"
)

var _ oas.Handler = (*handler)(nil)

type handler struct {
	oas.UnimplementedHandler
	queries *db.Queries
}

func NewHandler(queries *db.Queries) *handler {
	return &handler{
		queries: queries,
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
