package helper

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func Json(w http.ResponseWriter, statusCode int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		return
	}
}

func JsonSuccess(w http.ResponseWriter, statusCode int, data any) {
	Json(w, statusCode, APIResponse{
		Success: true,
		Data:    data,
	})
}

func JsonMessage(w http.ResponseWriter, statusCode int, message string) {
	Json(w, statusCode, APIResponse{
		Success: true,
		Message: message,
	})
}

func JsonError(w http.ResponseWriter, statusCode int, errorMessage string) {
	Json(w, statusCode, APIResponse{
		Success: false,
		Error:   errorMessage,
	})
}
