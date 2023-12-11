package geocode

import (
	"sync"
	"time"
)

type RateLimit struct {
	slot  int64
	hits  int
	Max   int
	mutex sync.Mutex
}

func (r *RateLimit) Claim() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	now := time.Now().Unix()
	if r.slot != now {
		logger.Debug("New slot created")
		r.slot = now
		r.hits = 1
		return true
	}
	r.hits++
	if r.hits > r.Max {
		logger.Debug("Rate limit hit")
		return false
	}
	logger.Debug("Rate limit incremented")
	return true
}
