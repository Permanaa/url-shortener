package handler

import (
	"github.com/permanaa/url-shortener/internal/dbgen"
)

type Handler struct {
	queries *dbgen.Queries
}

func NewHandler(queries *dbgen.Queries) *Handler {
	return &Handler{
		queries: queries,
	}
}
