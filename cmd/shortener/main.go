package main

import (
	"log"
	"net/http"
	"os"

	"github.com/f1atee/url-shortener/internal/handler"
	"github.com/f1atee/url-shortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://shortener:shortener@localhost:5432/shortener?sslmode=disable"
	}

	store, err := storage.New(dsn)
	if err != nil {
		log.Fatalf("failed connect to db: %v", err)
	}
	defer store.Close()

	h := handler.New(store)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/api/shorten", h.Shorten)
	r.Get("/{code}", h.Redirect)
	r.Delete("/{code}", h.Delete)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
