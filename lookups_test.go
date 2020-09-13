package lookups_test

import (
	"testing"

	"github.com/4lie/lookups"
	"github.com/golang/geo/s2"
	"github.com/stretchr/testify/assert"
)

func TestLookups(t *testing.T) {
	a := assert.New(t)

	rawPolygon := [][]float64{
		{
			35.837037430733666,
			50.993492603302,
		},
		{
			35.837037430733666,
			50.994919538497925,
		},
		{
			35.837837616204105,
			50.994919538497925,
		},
		{
			35.837837616204105,
			50.993492603302,
		},
	}

	points := make([]s2.Point, 0, len(rawPolygon))

	for _, p := range rawPolygon {
		points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(p[0], p[1])))
	}

	polygon := s2.PolygonFromLoops([]*s2.Loop{s2.LoopFromPoints(points)})

	l := lookups.NewWithS2([]lookups.PolyProps{
		{
			Props: lookups.Props{
				"score": 20,
			},
			Polygon: polygon,
		},
	}, lookups.DefaultS2CellLevel)

	r := l.Lookup([]lookups.Coordinate{
		{
			Latitude:  35.83738316403,
			Longitude: 50.99430799484,
		},
		{
			Latitude:  35.83599369842,
			Longitude: 51.00250482559,
		},
	})

	a.Len(r, 2)

	a.Len(r[0].Props, 1)
	s, ok := r[0].Props[0]["score"].(int)
	a.True(ok)
	a.Equal(20, s)

	a.Empty(r[1].Props)
}
