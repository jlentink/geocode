package main

import (
	"fmt"
	"github.com/jlentink/geocode"
	log "github.com/jlentink/yaglogger"
)

func main() {
	geo := geocode.NewGeoCode(geocode.WithLogger(log.GetInstance()))
	loc, err := geo.EncodeParametrized(geocode.FindParam{Country: "Netherlands", City: "Den Haag", Street: "'s-Gravenhaagse Bos"})
	if err != nil {
		log.Error("err: %v", err)
		return
	}
	for _, location := range loc.Locations {
		fmt.Printf("%s\n", location.DisplayName)
	}
}
