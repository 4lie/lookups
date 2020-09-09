package main

import (
	"log"

	"github.com/4lie/lookups"
)

func main() {
	_, err := lookups.NewPipEngine(nil)
	if err != nil {
		log.Fatalf("failed to create lookups point in polygon engine: %s", err.Error())
	}
}
