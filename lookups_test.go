package lookups_test

import (
	"testing"

	"github.com/4lie/lookups"
	"github.com/golang/geo/s2"
	"github.com/stretchr/testify/assert"
)

//nolint: funlen
func TestLookups(t *testing.T) {
	tests := []struct {
		name           string
		polyProps      []lookups.PolyProps
		queries        []lookups.Coordinate
		expectedResult [][]lookups.Props
	}{
		{
			name: "simple polygon",
			polyProps: []lookups.PolyProps{
				polyPropsFromCoordinates([][]lookups.Coordinate{
					{
						{
							Latitude:  35.837037430733666,
							Longitude: 50.993492603302,
						},
						{
							Latitude:  35.837037430733666,
							Longitude: 50.994919538497925,
						},
						{
							Latitude:  35.837837616204105,
							Longitude: 50.994919538497925,
						},
						{
							Latitude:  35.837837616204105,
							Longitude: 50.993492603302,
						},
					},
				},
					map[string]interface{}{
						"polygon": "blue",
					},
				),
			},
			queries: []lookups.Coordinate{
				{
					Latitude:  35.837411,
					Longitude: 50.994426,
				},
				{
					Latitude:  35.836411,
					Longitude: 50.996733,
				},
			},
			expectedResult: [][]lookups.Props{
				{
					{"polygon": "blue"},
				},
				{},
			},
		},
		{
			name: "polygon with holes",
			polyProps: []lookups.PolyProps{
				polyPropsFromCoordinates([][]lookups.Coordinate{
					{
						{
							Latitude:  35.839542,
							Longitude: 50.991776,
						},
						{
							Latitude:  35.834498,
							Longitude: 50.991669,
						},
						{
							Latitude:  35.834289,
							Longitude: 51.002011,
						},
						{
							Latitude:  35.83956,
							Longitude: 51.001968,
						},
					},
					{
						{
							Latitude:  35.839194,
							Longitude: 50.993954,
						},
						{
							Latitude:  35.837316,
							Longitude: 50.992484,
						},
						{
							Latitude:  35.836977,
							Longitude: 50.996207,
						},
					},
					{
						{
							Latitude:  35.838012,
							Longitude: 50.999587,
						},
						{
							Latitude:  35.835472,
							Longitude: 50.997097,
						},
						{
							Latitude:  35.835141,
							Longitude: 51.001368,
						},
					},
				},
					map[string]interface{}{
						"polygon": "blue",
					},
				),
			},
			queries: []lookups.Coordinate{
				{
					Latitude:  35.840743,
					Longitude: 50.996776,
				},
				{
					Latitude:  35.837681,
					Longitude: 50.994222,
				},
				{
					Latitude:  35.835976,
					Longitude: 50.999308,
				},
				{
					Latitude:  35.838429,
					Longitude: 51.000767,
				},
			},
			expectedResult: [][]lookups.Props{
				{},
				{},
				{},
				{
					{"polygon": "blue"},
				},
			},
		},
		{
			name: "polygon with holes and a simple polygon",
			polyProps: []lookups.PolyProps{
				polyPropsFromCoordinates([][]lookups.Coordinate{
					{
						{
							Latitude:  35.839542,
							Longitude: 50.991776,
						},
						{
							Latitude:  35.834498,
							Longitude: 50.991669,
						},
						{
							Latitude:  35.834289,
							Longitude: 51.002011,
						},
						{
							Latitude:  35.83956,
							Longitude: 51.001968,
						},
					},
					{
						{
							Latitude:  35.839194,
							Longitude: 50.993954,
						},
						{
							Latitude:  35.837316,
							Longitude: 50.992484,
						},
						{
							Latitude:  35.836977,
							Longitude: 50.996207,
						},
					},
					{
						{
							Latitude:  35.838012,
							Longitude: 50.999587,
						},
						{
							Latitude:  35.835472,
							Longitude: 50.997097,
						},
						{
							Latitude:  35.835141,
							Longitude: 51.001368,
						},
					},
				},

					map[string]interface{}{
						"polygon": "blue",
					},
				),
				polyPropsFromCoordinates([][]lookups.Coordinate{
					{
						{
							Latitude:  35.84163,
							Longitude: 50.995038,
						},
						{
							Latitude:  35.837455,
							Longitude: 51.001432,
						},
						{
							Latitude:  35.841925,
							Longitude: 51.001239,
						},
					},
				},

					map[string]interface{}{
						"polygon": "green",
					},
				),
			},
			queries: []lookups.Coordinate{
				{
					Latitude:  35.840743,
					Longitude: 50.996776,
				},
				{
					Latitude:  35.837681,
					Longitude: 50.994222,
				},
				{
					Latitude:  35.835976,
					Longitude: 50.999308,
				},
				{
					Latitude:  35.838429,
					Longitude: 51.000767,
				},
				{
					Latitude:  35.835524,
					Longitude: 50.993471,
				},
			},
			expectedResult: [][]lookups.Props{
				{
					{"polygon": "green"},
				},
				{},
				{},
				{
					{"polygon": "blue"},
					{"polygon": "green"},
				},
				{
					{"polygon": "blue"},
				},
			},
		},
	}
	for i := range tests {
		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			a := assert.New(t)

			a.Equal(len(test.expectedResult), len(test.queries))

			l := lookups.New(test.polyProps, lookups.NewS2Index(lookups.DefaultS2CellLevel))

			result := l.Lookup(test.queries)

			a.Equal(len(test.queries), len(result))

			for i := range result {
				a.Equal(test.expectedResult[i], result[i].Props)
			}
		})
	}
}

func polyPropsFromCoordinates(in [][]lookups.Coordinate, props map[string]interface{}) lookups.PolyProps {
	return lookups.PolyProps{
		Polygon: polygonFromCoordinates(in),
		Props:   props,
	}
}

func polygonFromCoordinates(in [][]lookups.Coordinate) *s2.Polygon {
	loops := make([]*s2.Loop, 0, len(in))

	for _, l := range in {
		loop := loopFromCoordinates(l)
		loop.Normalize()
		loops = append(loops, loop)
	}

	return s2.PolygonFromLoops(loops)
}

func loopFromCoordinates(in []lookups.Coordinate) *s2.Loop {
	points := make([]s2.Point, 0, len(in))

	for _, p := range in {
		points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(p.Latitude, p.Longitude)))
	}

	return s2.LoopFromPoints(points)
}
