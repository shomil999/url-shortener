package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/shomil999/url-shortener/internal/metrics"
	"github.com/shomil999/url-shortener/internal/shortener"
)

type Handlers struct {
	svc *shortener.Service
	met *metrics.Metrics
}

func NewHandlers(svc *shortener.Service, met *metrics.Metrics) *Handlers {
	return &Handlers{svc: svc, met: met}
}

//Shorten URL handler
func (h *Handlers) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req struct{ URL string `json:"url"` }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || strings.TrimSpace(req.URL) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}
	short, code, err := h.svc.Shorten(req.URL)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"short_url": short, "code": code})
}

//Redirect handler
func (h *Handlers) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" || strings.HasPrefix(code, "api/") {
		if code == "" {
			writeJSON(w, http.StatusOK, map[string]string{"message": "Go URL Shortener Service"})
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}
	url, err := h.svc.Resolve(code)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("code not found"))
		return
	}
	log.Printf("Redirecting %s -> %s", code, url)
	http.Redirect(w, r, url, http.StatusFound)
}

//Top Domains handler
func (h *Handlers) TopDomains(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	writeJSON(w, http.StatusOK, h.met.TopN(3))
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}