package geocode

import (
	"strings"
)

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		cache: make(map[string][]*Location),
	}
}

type InMemoryCache struct {
	cache map[string][]*Location
}

func (i InMemoryCache) NormalizeKey(key string) string {
	key = strings.ToLower(key)
	key = strings.TrimSpace(key)
	return key
}

func (i InMemoryCache) Exists(key string) bool {
	_, ok := i.cache[i.NormalizeKey(key)]
	return ok
}

func (i InMemoryCache) Get(key string) (*Response, error) {
	l, ok := i.cache[i.NormalizeKey(key)]

	if !ok {
		return NewResponse(), ErrNotFound
	}
	r := NewResponse()
	r.Locations = l
	r.Cached = true
	r.RetryAfter = 0
	return r, nil
}

func (i InMemoryCache) Set(key string, value []*Location) error {
	i.cache[i.NormalizeKey(key)] = value
	return nil
}
