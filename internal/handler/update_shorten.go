package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/permanaa/url-shortener/internal/dbgen"
	"github.com/permanaa/url-shortener/internal/helper"
)

type UpdateShortenRequest struct {
	Url string `json:"url"`
}

func (h *Handler) UpdateShorten(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")

	var req UpdateShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("err decode req Body update shorten: ", err)
		helper.JsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	findUrl, err := h.queries.GetUrlByShortCode(r.Context(), shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helper.JsonError(w, http.StatusNotFound, "Url doesn't exist")
			return
		}
		log.Println("err find url: ", err)
		helper.JsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	validatedUrl, err := helper.ValidateURL(req.Url)
	if err != nil {
		helper.JsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	updateUrl, err := h.queries.UpdateUrlByShortCode(r.Context(), dbgen.UpdateUrlByShortCodeParams{
		Url:       validatedUrl,
		ShortCode: findUrl.ShortCode,
	})
	if err != nil {
		log.Println("err update url: ", err)
		helper.JsonError(w, http.StatusInternalServerError, "Failed to update url")
		return
	}

	urlRes := UrlResponse{
		ID:        updateUrl.ID,
		CreatedAt: updateUrl.CreatedAt,
		UpdatedAt: updateUrl.UpdatedAt,
		Url:       updateUrl.Url,
		ShortCode: updateUrl.ShortCode,
	}

	helper.JsonSuccess(w, http.StatusOK, urlRes)
}
