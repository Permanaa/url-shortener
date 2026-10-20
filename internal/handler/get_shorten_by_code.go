package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/permanaa/url-shortener/internal/helper"
)

func (h *Handler) GetShortenByCode(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")

	findUrl, err := h.queries.GetUrlByShortCode(r.Context(), shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helper.JsonError(w, http.StatusNotFound, "Url not found")
			return
		}

		helper.JsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	helper.JsonSuccess(w, http.StatusOK, findUrl)
}
