package geocode

import (
	"net/url"
	"testing"
)

func TestFindParam_ToString(t *testing.T) {
	param := FindParam{
		Street:     "123 Main St",
		City:       "Anytown",
		County:     "Anycounty",
		State:      "Anystate",
		Country:    "Anycountry",
		PostalCode: "12345",
	}

	expected := "123 Main St, Anytown, Anycounty, Anystate, Anycountry, 12345"
	if str := param.ToString(); str != expected {
		t.Errorf("Expected '%s', got '%s'", expected, str)
	}
}

func TestFindParam_URLValues(t *testing.T) {
	param := FindParam{
		Street:     "123 Main St",
		City:       "Anytown",
		County:     "Anycounty",
		State:      "Anystate",
		Country:    "Anycountry",
		PostalCode: "12345",
	}

	expected := url.Values{
		QueryStreet:     []string{"123 Main St"},
		QueryCity:       []string{"Anytown"},
		QueryCounty:     []string{"Anycounty"},
		QueryState:      []string{"Anystate"},
		QueryCountry:    []string{"Anycountry"},
		QueryPostalCode: []string{"12345"},
	}

	if compareURLValues(param.URLValues(), expected) {
		t.Errorf("Expected '%s', got '%s'", expected, param.URLValues())
	}
}

func compareURLValues(values1, values2 url.Values) bool {
	if len(values1) != len(values2) {
		return false
	}

	for key, values := range values1 {
		if v, ok := values2[key]; !ok {
			return false
		} else {
			if !compareStringSlices(values, v) {
				return false
			}
		}
	}

	return true
}

func compareStringSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}
