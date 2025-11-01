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

var ErrInvalidURL = errors.New("invalid url: must be absolute http/https")

func (s *Service) Shorten(raw string) (shortURL, code string, err error) {
	u, err := url.Parse(raw)
	if err != nil || !u.IsAbs() || (u.Scheme != "http" && u.Scheme != "https") {
		return "", "", ErrInvalidURL
	}
	u.Host = strings.ToLower(u.Host)
	norm := u.String()

	if existing, err2 := s.store.GetByURL(norm); err2 == nil {
		s.metrics.IncDomain(util.DomainKey(u.Host))
		return s.baseURL + "/" + existing, existing, nil
	}

	sum := sha256.Sum256([]byte(norm))
	code = base62Encode(sum[:6])
	if err := s.store.Save(norm, code); err != nil {
		return "", "", err
	}

	s.metrics.IncDomain(util.DomainKey(u.Host))
	return s.baseURL + "/" + code, code, nil
}

func (s *Service) Resolve(code string) (string, error) {
	return s.store.GetByCode(code)
}
