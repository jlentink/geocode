package geocode

type ResponseOptFunc func(*ResponseOpt)

type ResponseOpt struct {
	Locations  []*Location
	Cached     bool
	RetryAfter int
}

func responseDefaultOpts() *ResponseOpt {
	return &ResponseOpt{
		Locations:  []*Location{},
		Cached:     false,
		RetryAfter: -1,
	}
}

func ResponseWithLocations(locations []*Location) ResponseOptFunc {
	return func(o *ResponseOpt) {
		o.Locations = locations
	}
}
func ResponseWithLocation(location *Location) ResponseOptFunc {
	return func(o *ResponseOpt) {
		o.Locations = []*Location{location}
	}
}

func ResponseWithCached(cached bool) ResponseOptFunc {
	return func(o *ResponseOpt) {
		o.Cached = cached
	}
}

func ResponseWithRetryAfter(retryAfter int) ResponseOptFunc {
	return func(o *ResponseOpt) {
		o.RetryAfter = retryAfter
	}
}
