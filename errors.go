package geocode

import "errors"

var ErrNotFound = errors.New("not found")
var ErrRateLimit = errors.New("rate limit exceeded")
var ErrBlocked = errors.New("your ip address is blocked")
var ErrServiceUnavailable = errors.New(" service not available")
