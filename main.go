package main

import (
	"flag"

	gen "github.com/Trashed/go-heights/hmapgen"
)

var (
	detail = flag.Int("detail", 6, `Generates height map based on given detail value with following
        formula: 2^d+1. Default detail value is 6 => array size 65.`)
	numOfMaps = flag.Int("amount", 1, "Number of maps to generate.")
)

func main() {
	flag.Parse()

	// Create new Terrain object and generate height map data.
	for i := 0; i < *numOfMaps; i++ {
		t := gen.New(*detail)
		t.Generate()

		// Create height map image.
		gen.Image(*t)
	}
}
