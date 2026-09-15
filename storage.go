package main

import (
	"sync"
)

type URLStore struct {
	mu   sync.Mutex
	urls map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{
		mu:   sync.Mutex{},
		urls: map[string]string{},
	}
}

func (s *URLStore) Add(shortCode, longURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.urls[shortCode] = longURL
	return nil
}

func (s *URLStore) Get(shortCode string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	urlCode, exists := s.urls[shortCode]
	return urlCode, exists
}

func (s *URLStore) All() map[string]string {
	s.mu.Lock()
	defer s.mu.Unlock()

	copy := make(map[string]string)
	for code, url := range s.urls {
		copy[code] = url
	}

	return copy
}

func (s *URLStore) Delete(shortCode string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.urls, shortCode)
}
