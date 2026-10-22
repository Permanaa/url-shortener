package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/permanaa/url-shortener/internal/dbgen"
	"github.com/permanaa/url-shortener/internal/handler"
	"github.com/permanaa/url-shortener/internal/helper"
	"github.com/permanaa/url-shortener/internal/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Load env: %v", err)
	}

	dbUrl := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Open db: %v", err)
	}
	dbQueries := dbgen.New(db)

	handler := handler.NewHandler(dbQueries)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		helper.JsonMessage(w, http.StatusOK, "Hello, Universe!")
	})
	mux.HandleFunc("POST /shortnd", handler.PostShortener)
	mux.HandleFunc("GET /shortnd/{shortCode}", handler.GetShortenByCode)
	mux.HandleFunc("PUT /shortnd/{shortCode}", handler.UpdateShorten)
	mux.HandleFunc("DELETE /shortnd/{shortCode}", handler.DeleteShorten)
	mux.HandleFunc("GET /shortnd/{shortCode}/stats", handler.GetStats)

	rateLimiter := middleware.NewRateLimiter()
	applyMiddleware := middleware.ApplyMiddleware(mux, rateLimiter.Apply)

	server := &http.Server{
		Addr:    ":1804",
		Handler: applyMiddleware,
	}

	log.Println("Serving on port: 1804")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server ListenAndServer: %v", err)
	}
}
