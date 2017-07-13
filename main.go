package main

import (
	"flag"
	"fmt"

	gen "github.com/Trashed/go-heights/hmapgen"
)

var (
	detail = flag.Int("detail", 6, `Generates height map based on given detail value with following
        formula: 2^d+1. Default detail value is 6 => array size 65.`)
)

func main() {
	flag.Parse()
	fmt.Printf("detail: %d\n", *detail)

	t := gen.New(*detail)
	t.Generate()
	fmt.Printf("terrain data: %v\n", t.Data())
}
