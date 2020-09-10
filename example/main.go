package main

import (
	"github.com/4lie/lookups"
)

func main() {
	lookups.NewPipEngineFromGeom(nil)
	lookups.NewPipEngineFromS2(nil)
}
