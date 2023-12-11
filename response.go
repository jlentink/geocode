package geocode

func NewResponse(opts ...ResponseOptFunc) *Response {
	o := responseDefaultOpts()
	for _, fn := range opts {
		fn(o)
	}
	return &Response{
		Locations:  o.Locations,
		Cached:     o.Cached,
		RetryAfter: o.RetryAfter,
	}
}

// Response is the wrapper for the geocode API response
// that contains the results, potential retry delay and cached status.
type Response struct {
	Locations  []*Location `json:"locations"`
	Cached     bool        `json:"cached"`
	RetryAfter int         `json:"retry_after"`
}

// Location is the response from the geocode API.
type Location struct {
	PlaceID     int64    `json:"place_id"`
	Licence     string   `json:"licence"`
	PoweredBy   string   `json:"powered_by"`
	OsmType     string   `json:"osm_type"`
	OsmID       int64    `json:"osm_id"`
	BoundingBox []string `json:"boundingbox"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
}
