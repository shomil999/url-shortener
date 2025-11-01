package shortener

import (
	"crypto/sha256"
	"errors"
	"net/url"
	"strings"

	"github.com/shomil999/url-shortener/internal/metrics"
	"github.com/shomil999/url-shortener/internal/store"
	"github.com/shomil999/url-shortener/pkg/util"
)

type Service struct {
	store   *store.MemoryStore
	metrics *metrics.Metrics
	baseURL string
}

func NewService(st *store.MemoryStore, m *metrics.Metrics, baseURL string) *Service {
	return &Service{store: st, metrics: m, baseURL: strings.TrimRight(baseURL, "/")}
}