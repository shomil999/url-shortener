package store

import (
	"errors"
	"sync"
)

var errNotFound = errors.New("not found")

type MemoryStore struct {
	mu        sync.RWMutex
	urlToCode map[string]string
	codeToURL map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		urlToCode: make(map[string]string),
		codeToURL: make(map[string]string),
	}
}

func (s *MemoryStore) GetByURL(url string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	code, ok := s.urlToCode[url]
	if !ok {
		return "", errNotFound
	}
	return code, nil
}

func (s *MemoryStore) GetByCode(code string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.codeToURL[code]
	if !ok {
		return "", errNotFound
	}
	return url, nil
}

func (s *MemoryStore) Save(url, code string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.urlToCode[url]; ok {
		return nil
	}
	s.urlToCode[url] = code
	s.codeToURL[code] = url
	return nil
}
