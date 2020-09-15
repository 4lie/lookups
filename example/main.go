package main

import (
	"fmt"

	"github.com/4lie/lookups"
)

//nolint:gomnd,funlen
func main() {
	polyProps := []lookups.PolyProps{
		lookups.PolyPropsFromCoordinates([][]lookups.Coordinate{
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
		lookups.PolyPropsFromCoordinates([][]lookups.Coordinate{
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
	}

	queries := []lookups.Coordinate{
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
	}

	lookuper := lookups.New(polyProps, lookups.DefaultGeoIndexer)

	coordinatesProps := lookuper.Lookup(queries)

	for _, coordinateProps := range coordinatesProps {
		fmt.Println(coordinateProps.Props)
	}
}
