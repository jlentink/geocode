package geocode

//[{"place_id":32279975,"licence":"Data © OpenStreetMap contributors, ODbL 1.0. https://osm.org/copyright","powered_by":"Map Maker: https://maps.co","osm_type":"node","osm_id":2812003297,"boundingbox":["52.5561276","52.5562276","4.6639743","4.6640743"],"lat":"52.5561776","lon":"4.6640243","display_name":"103, Iepenlaan, Castricum, North Holland, Netherlands, 1901SV, Netherlands","class":"place","type":"house","importance":0.31100000000000005}]

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

type Response struct {
	Locations  []*Location `json:"locations"`
	Cached     bool        `json:"cached"`
	RetryAfter int         `json:"retry_after"`
}

// Location is the response from the geocode API
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
