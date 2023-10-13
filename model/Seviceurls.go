package model

import (
	"ginLearnDemo/utils"
	"log"
	"sync"
)

type URLStore struct {
	Urls map[string]string
	mu   sync.RWMutex
}

func NewUrlStore() *URLStore {
	return &URLStore{Urls: make(map[string]string)}
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Urls[key]
}

func (s *URLStore) Set(key, value string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.Urls[key]
	log.Println(key, value, ok)
	if ok {
		return false
	}
	s.Urls[key] = value
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.Urls)
}

func (s *URLStore) Put(value string) string {
	for {
		key := utils.Base62()
		if s.Set(key, value) {
			return key
		}
	}
	return ""
}
