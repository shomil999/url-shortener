package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shomil999/url-shortener/internal/httpapi"
	"github.com/shomil999/url-shortener/internal/metrics"
	"github.com/shomil999/url-shortener/internal/shortener"
	"github.com/shomil999/url-shortener/internal/store"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:" + port
	}

	st := store.NewMemoryStore()
	met := metrics.New()
	svc := shortener.NewService(st, met, baseURL)
	h := httpapi.NewHandlers(svc, met)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/shorten", h.ShortenURL)
	mux.HandleFunc("/api/v1/metrics/top-domains", h.TopDomains)
	mux.HandleFunc("/", h.Redirect)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      loggingMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("URL Shortener running on %s (base=%s)", srv.Addr, baseURL)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
