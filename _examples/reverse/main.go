package main

import (
	"fmt"

	"github.com/jlentink/geocode"
)

func main() {
	geo := geocode.NewGeoCode()
	loc, err := geo.ReverseEncode(52.0915319, 4.3433845)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	for _, location := range loc.Locations {
		fmt.Printf("%s\n", location.DisplayName)
	}
}
