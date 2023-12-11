package main

import (
	"fmt"

	"github.com/jlentink/geocode"
)

func main() {
	geo := geocode.NewGeoCode()
	loc, err := geo.EncodeParametrized(geocode.FindParam{Country: "Netherlands", City: "Den Haag", Street: "'s-Gravenhaagse Bos"})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	for _, location := range loc.Locations {
		fmt.Printf("%s\n", location.DisplayName)
	}
}
