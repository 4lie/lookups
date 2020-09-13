package index_test

import (
	"testing"

	"github.com/4lie/lookups/index"

	"github.com/golang/geo/s2"
	"github.com/stretchr/testify/assert"
)

func TestS2Hash(t *testing.T) {
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

	s := index.NewS2(15)

	ids := s.Cover(polygon)

	a.Len(ids, 2)
	a.Contains(ids, "3f8dbf754")
	a.Contains(ids, "3f8dbf75c")
}
