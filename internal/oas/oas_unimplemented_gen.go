// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateShortURL implements createShortURL operation.
//
// Create a shorten URL.
//
// POST /create
func (UnimplementedHandler) CreateShortURL(ctx context.Context, req *CreateShortURLReq) (r CreateShortURLRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RedirectLongURL implements redirectLongURL operation.
//
// Redirect client to long URL.
//
// GET /{hash}
func (UnimplementedHandler) RedirectLongURL(ctx context.Context, params RedirectLongURLParams) (r RedirectLongURLRes, _ error) {
	return r, ht.ErrNotImplemented
}
