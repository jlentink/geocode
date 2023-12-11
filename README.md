# Geocode

This is a simple Go package for geocoding addresses using the [Geocode API](https://geocode.maps.co/).

## Installation

```bash
go get github.com/jlentink/geocode
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/jlentink/geocode"
)

func main() {
    address := "'s-Gravenhaagse Bos, Den haag, Netherlands"
    lat, lon, err := geocode.Geocode(address)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Latitude: %f, Longitude: %f\n", lat, lon)
}
```
All other example can be found in the [examples](_examples) directory.

## Acknowledgments

* Thanks to [Geocode API](https://geocode.maps.co/) for their service.