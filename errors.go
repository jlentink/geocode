package geocode

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrRateLimit          = errors.New("rate limit exceeded")
	ErrBlocked            = errors.New("your ip address is blocked")
	ErrServiceUnavailable = errors.New(" service not available")
	//goland:noinspection
	ErrCacheFailed = errors.New("cache failed")
)
