package main

import (
	"fmt"

	"github.com/jlentink/geocode"
)

func main() {
	geo := geocode.NewGeoCode()
	loc, err := geo.Encode("'s-Gravenhaagse Bos, Den haag, Netherlands")
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	for _, location := range loc.Locations {
		fmt.Printf("%s\n", location.DisplayName)
	}
}
