package URLCache

import (
	"sync"
)

type URLCache struct {
	mu  sync.Mutex
	Cache map[string]bool
}

func (s *URLCache) Visit(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Cache[url] {
		return true
	}
	s.Cache[url] = true
	return false
}