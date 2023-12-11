package geocode

type OptFunc func(*Opts)

type Opts struct {
	cache            ICache
	rateLimit        int
	searchURL        string
	reverseSearchURL string
	log              Loggable
}

// WithCache sets the cache for the geocode service.
//
//goland:noinspection GoUnusedExportedFunction
func WithCache(cache ICache) OptFunc {
	return func(o *Opts) {
		o.cache = cache
	}
}

// WithRateLimit sets the rate limit for the geocode service.
//
//goland:noinspection GoUnusedExportedFunction
func WithRateLimit(rateLimit int) OptFunc {
	return func(o *Opts) {
		o.rateLimit = rateLimit
	}
}

// WithLogger sets the logger for the geocode service.
//
//goland:noinspection GoUnusedExportedFunction
func WithLogger(log Loggable) OptFunc {
	return func(o *Opts) {
		o.log = log
	}
}

func defaultOpts() *Opts {
	return &Opts{
		cache:            NewInMemoryCache(),
		log:              NoLog{},
		rateLimit:        1,
		searchURL:        "https://geocode.maps.co/search",
		reverseSearchURL: "https://geocode.maps.co/reverse",
	}
}
