package lookups

import (
	"github.com/twpayne/go-geom/encoding/wkb"
)

type (
	Coordinate struct {
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
		SequenceID int     `json:"sequence_id"`
	}

	PolyProps struct {
		Polygon *wkb.Polygon           `json:"polygon"`
		Props   map[string]interface{} `json:"props"`
	}

	CoordinateProps struct {
		Coordinate Coordinate `json:"coordinate"`
		PolyProps  PolyProps  `json:"poly_props"`
	}
)
