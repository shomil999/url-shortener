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