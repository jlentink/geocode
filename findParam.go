package geocode

import "net/url"

const (
	QueryStreet     = "street"
	QueryCity       = "city"
	QueryCounty     = "county"
	QueryState      = "state"
	QueryCountry    = "country"
	QueryPostalCode = "postalCode"
)

type FindParam struct {
	Street     string
	City       string
	County     string
	State      string
	Country    string
	PostalCode string
}

func (f FindParam) ToString() string {
	return f.Street + ", " + f.City + ", " + f.County + ", " + f.State + ", " + f.Country + ", " + f.PostalCode
}

func (f FindParam) URLValues() url.Values {
	v := url.Values{}
	if f.Street != "" {
		v.Add(QueryStreet, f.Street)
	}
	if f.City != "" {
		v.Add(QueryCity, f.City)
	}
	if f.County != "" {
		v.Add(QueryCounty, f.County)
	}
	if f.State != "" {
		v.Add(QueryState, f.State)
	}
	if f.Country != "" {
		v.Add(QueryCountry, f.State)
	}
	if f.PostalCode != "" {
		v.Add(QueryPostalCode, f.PostalCode)
	}
	return v
}
