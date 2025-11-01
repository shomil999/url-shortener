package store

import "testing"

func TestMemoryStore_SaveAndGet(t *testing.T) {
	st := NewMemoryStore()
	url := "https://example.com/test"
	code := "abc123"

	// Save mapping
	if err := st.Save(url, code); err != nil {
		t.Fatalf("unexpected error while saving: %v", err)
	}

	gotCode, err := st.GetByURL(url)
	if err != nil {
		t.Fatalf("expected to find code for URL %q, got error: %v", url, err)
	}
	if gotCode != code {
		t.Errorf("expected code %q, got %q", code, gotCode)
	}

	gotURL, err := st.GetByCode(code)
	if err != nil {
		t.Fatalf("expected to find URL for code %q, got error: %v", code, err)
	}
	if gotURL != url {
		t.Errorf("expected URL %q, got %q", url, gotURL)
	}
}

func TestMemoryStore_IdempotentSave(t *testing.T) {
	st := NewMemoryStore()
	url := "https://example.com/idempotent"
	code1 := "code1"
	code2 := "code2"

	if err := st.Save(url, code1); err != nil {
		t.Fatalf("unexpected error while saving: %v", err)
	}
	if err := st.Save(url, code2); err != nil {
		t.Fatalf("unexpected error while saving: %v", err)
	}

	got, err := st.GetByURL(url)
	if err != nil {
		t.Fatalf("expected to find URL, got error: %v", err)
	}
	if got != code1 {
		t.Errorf("expected original code %q to remain, got %q", code1, got)
	}
}
