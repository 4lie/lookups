package lookups

import (
	"github.com/twpayne/go-geom/encoding/wkb"
)

type (
	Props map[string]interface{}

	Coordinate struct {
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
		SequenceID int     `json:"sequence_id"`
	}

	PolyProps struct {
		Props   Props        `json:"props"`
		Polygon *wkb.Polygon `json:"polygon"`
	}

	CoordinateProps struct {
		Props      []Props    `json:"props"`
		Coordinate Coordinate `json:"coordinate"`
	}
)
