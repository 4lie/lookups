package lookups

import (
	"github.com/golang/geo/s2"
	"github.com/twpayne/go-geom"
)

type (
	Props map[string]interface{}

	Coordinate struct {
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
		SequenceID int     `json:"sequence_id"`
	}

	GeomPolyProps struct {
		Props   Props         `json:"props"`
		Polygon *geom.Polygon `json:"polygon"`
	}

	S2PolyProps struct {
		Props   Props       `json:"props"`
		Polygon *s2.Polygon `json:"polygon"`
	}

	CoordinateProps struct {
		Props      []Props    `json:"props"`
		Coordinate Coordinate `json:"coordinate"`
	}
)
