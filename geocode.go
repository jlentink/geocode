package geocode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var logger Loggable = NoLog{}

const (
	QueryGeneral = "q"
	QueryLat     = "lat"
	QueryLon     = "lon"
)

// NewGeoCode creates a new GeoCode struct with the provided options.
//
//goland:noinspection GoUnusedExportedFunction
func NewGeoCode(opts ...OptFunc) *GeoCode {
	o := defaultOpts()
	for _, fn := range opts {
		fn(o)
	}

	logger = o.log
	return &GeoCode{
		options: o,
		cache:   o.cache,
		rateLimit: &RateLimit{
			Max: o.rateLimit,
		},
	}
}

// GeoCode is the main struct for the geocode package
// that works with the https://geocode.maps.co/ API.
type GeoCode struct {
	options   *Opts
	cache     ICache
	rateLimit *RateLimit
}

// Encode takes a string and returns a Response struct with the results
// and cached status.
//
// When retry is set to anything else than -1 this means the service
// requested for a delay before the next request can be made.
func (g *GeoCode) Encode(subject string) (*Response, error) {
	if g.cache.Exists(subject) {
		resp, err := g.cache.Get(subject)
		return resp, err
	}
	if !g.rateLimit.Claim() {
		r := NewResponse()
		r.RetryAfter = 1
		return r, ErrRateLimit
	}
	resp, delay, err := g.httpReq(g.options.searchURL, url.Values{QueryGeneral: {subject}})
	if err != nil {
		return NewResponse(ResponseWithRetryAfter(delay)), err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var locations []*Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return nil, err
	}
	e := g.cache.Set(subject, locations)
	if e != nil {
		g.options.log.Error("Could not persist cache: %s", e)
	}
	return NewResponse(ResponseWithLocations(locations), ResponseWithCached(false)), nil
}

// EncodeParametrized takes a FindParam struct and returns a Response struct
// for more specific searching.
func (g *GeoCode) EncodeParametrized(param FindParam) (*Response, error) {
	subject := param.ToString()
	if g.cache.Exists(subject) {
		logger.Debug("Cache hit for %s", subject)
		return g.cache.Get(subject)
	}
	if !g.rateLimit.Claim() {
		logger.Debug("Rate limit hit for %s", subject)
		return nil, ErrRateLimit
	}

	resp, delay, err := g.httpReq(g.options.searchURL, param.URLValues())
	if err != nil {
		return NewResponse(ResponseWithRetryAfter(delay)), err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	var locations []*Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return nil, err
	}

	return NewResponse(ResponseWithLocations(locations), ResponseWithCached(false)), nil
}

// ReverseEncode takes a latitude and longitude and returns a Response struct.
func (g *GeoCode) ReverseEncode(lat, log float64) (*Response, error) {
	sLat := fmt.Sprintf("%f", lat)
	sLon := fmt.Sprintf("%f", log)
	subject := sLat + "," + sLon
	if g.cache.Exists(subject) {
		r, err := g.cache.Get(subject)
		return r, err
	}
	if !g.rateLimit.Claim() {
		return nil, ErrRateLimit
	}

	resp, delay, err := g.httpReq(g.options.reverseSearchURL, url.Values{QueryLat: {sLat}, QueryLon: {sLon}})
	if err != nil {
		return NewResponse(ResponseWithRetryAfter(delay)), err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var location *Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		return nil, err
	}

	return NewResponse(ResponseWithLocation(location), ResponseWithCached(false)), nil
}

// httpReq is a helper function that makes the actual request to the API.
func (g *GeoCode) httpReq(uri string, values url.Values) (*http.Response, int, error) {
	var retryAfter int
	uri = uri + "?" + values.Encode()

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, 1, err
	}

	resp, err := client.Do(req)
	if err != nil {
		if resp.StatusCode == http.StatusTooManyRequests {
			err = ErrRateLimit
			retryAfter = 1
		}
		if resp.StatusCode == http.StatusServiceUnavailable {
			err = ErrServiceUnavailable
			delay := resp.Header.Get("Retry-After")
			if delay != "" {
				d, err := strconv.Atoi(delay)
				if err == nil {
					retryAfter = d
				}
			}
		}
		if resp.StatusCode == http.StatusForbidden {
			err = ErrBlocked
		}
	}
	return resp, retryAfter, err
}
