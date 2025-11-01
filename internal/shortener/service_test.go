package shortener

import (
	"testing"

	"github.com/shomil999/url-shortener/internal/metrics"
	"github.com/shomil999/url-shortener/internal/store"
)

func TestShorten_Idempotent(t *testing.T) {
	st := store.NewMemoryStore()
	met := metrics.New()
	svc := NewService(st, met, "http://localhost:8080")

	url := "https://example.com/same"
	short1, code1, err := svc.Shorten(url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	short2, code2, err := svc.Shorten(url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if code1 != code2 {
		t.Errorf("expected same code for same URL, got %q and %q", code1, code2)
	}

	if short1 != short2 {
		t.Errorf("expected same short URLs, got %q and %q", short1, short2)
	}
}

func TestShorten_DifferentURLs(t *testing.T) {
	st := store.NewMemoryStore()
	met := metrics.New()
	svc := NewService(st, met, "http://localhost:8080")

	url1 := "https://google.com/abc"
	url2 := "https://google.com/xyz"

	_, code1, _ := svc.Shorten(url1)
	_, code2, _ := svc.Shorten(url2)

	if code1 == code2 {
		t.Errorf("expected different codes for different URLs, got same %q", code1)
	}
}

func TestResolve(t *testing.T) {
	st := store.NewMemoryStore()
	met := metrics.New()
	svc := NewService(st, met, "http://localhost:8080")

	url := "https://example.com/test"
	_, code, _ := svc.Shorten(url)

	got, err := svc.Resolve(code)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got != url {
		t.Errorf("expected URL %q, got %q", url, got)
	}
}
