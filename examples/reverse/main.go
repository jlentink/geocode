package main

import (
	"fmt"
	"github.com/jlentink/geocode"
	log "github.com/jlentink/yaglogger"
)

func main() {
	geo := geocode.NewGeoCode(geocode.WithLogger(log.GetInstance()))
	loc, err := geo.ReverseEncode(52.0915319, 4.3433845)
	if err != nil {
		log.Error("err: %v", err)
		return
	}
	for _, location := range loc.Locations {
		fmt.Printf("%s\n", location.DisplayName)
	}
}
