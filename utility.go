package lookups

import "github.com/golang/geo/s2"

func PolyPropsFromCoordinates(coordinates [][]Coordinate, props map[string]interface{}) PolyProps {
	return PolyProps{
		Polygon: PolygonFromCoordinates(coordinates),
		Props:   props,
	}
}

func PolygonFromCoordinates(coordinates [][]Coordinate) *s2.Polygon {
	loops := make([]*s2.Loop, 0, len(coordinates))

	for _, l := range coordinates {
		loop := LoopFromCoordinates(l)
		loop.Normalize()
		loops = append(loops, loop)
	}

	return s2.PolygonFromLoops(loops)
}

func LoopFromCoordinates(coordinates []Coordinate) *s2.Loop {
	points := make([]s2.Point, 0, len(coordinates))

	for _, p := range coordinates {
		points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(p.Latitude, p.Longitude)))
	}

	return s2.LoopFromPoints(points)
}
