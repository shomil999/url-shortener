package metrics

import "testing"

func TestMetrics_IncDomainAndTopN(t *testing.T) {
	m := New()

	// simulate counts
	m.IncDomain("udemy.com")
	m.IncDomain("udemy.com")
	m.IncDomain("youtube.com")
	m.IncDomain("wikipedia.org")
	m.IncDomain("wikipedia.org")
	m.IncDomain("wikipedia.org")

	top := m.TopN(3)

	if len(top) != 3 {
		t.Fatalf("expected 3 top domains, got %d", len(top))
	}

	if top[0].Domain != "wikipedia.org" {
		t.Errorf("expected top domain wikipedia.org, got %s", top[0].Domain)
	}

	if top[0].Count != 3 {
		t.Errorf("expected top count 3, got %d", top[0].Count)
	}
}
