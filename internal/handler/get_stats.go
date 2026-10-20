package handler

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/permanaa/url-shortener/internal/helper"
)

type StatsResponse struct {
	ID          uuid.UUID     `json:"id"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Url         string        `json:"url"`
	ShortCode   string        `json:"shortCode"`
	AccessCount sql.NullInt64 `json:"accessCount"`
}

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")

	findUrl, err := h.queries.GetUrlByShortCode(r.Context(), shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helper.JsonError(w, http.StatusNotFound, "Url not found")
			return
		}

		log.Println("err find url by short code:", err)
		helper.JsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	urlRes := StatsResponse{
		ID:          findUrl.ID,
		CreatedAt:   findUrl.CreatedAt,
		UpdatedAt:   findUrl.UpdatedAt,
		Url:         findUrl.Url,
		ShortCode:   findUrl.ShortCode,
		AccessCount: findUrl.AccessCount,
	}

	helper.JsonSuccess(w, http.StatusOK, urlRes)
}
