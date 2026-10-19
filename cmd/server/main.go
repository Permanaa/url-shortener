package main

import (
	"log"
	"net/http"

	"github.com/permanaa/url-shortener/internal/helper"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		helper.JsonMessage(w, http.StatusOK, "Hello, Universe!")
	})

	server := &http.Server{
		Addr:    ":1804",
		Handler: mux,
	}

	log.Println("Serving on port: 1804")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server ListenAndServer: %v", err)
	}
}
