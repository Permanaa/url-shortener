package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/permanaa/url-shortener/internal/dbgen"
	"github.com/permanaa/url-shortener/internal/helper"
)

func (h *Handler) PostShortener(w http.ResponseWriter, r *http.Request) {
	var req dbgen.CreateUrlParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("err decode req", err)
		helper.JsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	var shortCode string

	if req.ShortCode != "" {
		findUrl, err := h.queries.GetUrlByShortCode(r.Context(), req.ShortCode)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.Println("err find url", err)
			helper.JsonError(w, http.StatusInternalServerError, "Something went wrong")
			return
		}
		if findUrl.ShortCode != "" {
			helper.JsonError(w, http.StatusConflict, "Short code already exists")
			return
		}
		shortCode = strings.TrimSpace(req.ShortCode)
		if len(shortCode) > 20 || len(shortCode) < 6 {
			helper.JsonError(w, http.StatusBadRequest, "Short code must be between 6 and 20 characters")
			return
		}
		if !regexp.MustCompile(`^[a-zA-Z0-9_-]{6,20}$`).MatchString(shortCode) {
			helper.JsonError(w, http.StatusBadRequest, "Use only letters, numbers, hyphens (-), and underscores (_)")
			return
		}
	} else {
		// [TODO] url shorten algorithm
		helper.JsonMessage(w, http.StatusOK, "[TODO] Auto generated short code")
		return
	}

	validatedUrl, err := helper.ValidateURL(req.Url)
	if err != nil {
		helper.JsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	urlCreated, err := h.queries.CreateUrl(r.Context(), dbgen.CreateUrlParams{
		Url:       validatedUrl,
		ShortCode: shortCode,
	})
	if err != nil {
		log.Println("err create url", err)
		helper.JsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	helper.JsonSuccess(w, http.StatusCreated, urlCreated)
}
