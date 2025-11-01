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