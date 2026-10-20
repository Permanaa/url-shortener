package handler

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/permanaa/url-shortener/internal/helper"
)

func (h *Handler) DeleteShorten(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")

	findUrl, err := h.queries.GetUrlByShortCode(r.Context(), shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helper.JsonError(w, http.StatusNotFound, "Url doesn't exist")
			return
		}
		log.Println("err find url in delete:", err)
		helper.JsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if err := h.queries.DeleteUrlByShortCode(r.Context(), findUrl.ShortCode); err != nil {
		helper.JsonError(w, http.StatusInternalServerError, "Failed to delete url")
		return
	}

	helper.JsonMessage(w, http.StatusOK, "Url successfully deleted")
}
