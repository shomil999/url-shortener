package metrics

import (
	"sort"
	"strings"
	"sync"
)

type Metrics struct {
	mu       sync.RWMutex
	byDomain map[string]int
}

func New() *Metrics {
	return &Metrics{byDomain: make(map[string]int)}
}

func (m *Metrics) IncDomain(domain string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.byDomain[strings.ToLower(domain)]++
}

type DomainCount struct {
	Domain string `json:"domain"`
	Count  int    `json:"count"`
}

func (m *Metrics) TopN(n int) []DomainCount {
	m.mu.RLock()
	defer m.mu.RUnlock()
	arr := make([]DomainCount, 0, len(m.byDomain))
	for d, c := range m.byDomain {
		arr = append(arr, DomainCount{Domain: d, Count: c})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].Count == arr[j].Count {
			return arr[i].Domain < arr[j].Domain
		}
		return arr[i].Count > arr[j].Count
	})
	if n > len(arr) {
		n = len(arr)
	}
	return arr[:n]
}