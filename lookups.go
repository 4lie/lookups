package lookups

import (
	"github.com/golang/geo/s2"
)

const (
	// Default level of s2 cells.
	DefaultS2CellLevel = 15
)

type (
	// Lookuper is an interface for lookup services.
	Lookuper interface {
		Lookup(coordinates []Coordinate) []CoordinateProps
	}

	// Lookups is an engine of lookups.
	lookups struct {
		geoIndex    GeoIndex
		geoPolygons map[string][]PolyProps
	}
)

// NewWithS2 create a new Lookups instance using S2 geo index.
func NewWithS2(polyProps []PolyProps, s2Level int) Lookuper {
	geoIndex := NewS2Index(s2Level)

	items := make(map[string][]PolyProps)

	for _, polyProp := range polyProps {
		ids := geoIndex.Cover(polyProp.Polygon)

		for _, id := range ids {
			items[id] = append(items[id], polyProp)
		}
	}

	return &lookups{
		geoIndex:    geoIndex,
		geoPolygons: items,
	}
}

// Lookup returns list of properties of given coordinates.
func (l lookups) Lookup(coordinates []Coordinate) []CoordinateProps {
	result := make([]CoordinateProps, 0, len(coordinates))

	for _, coordinate := range coordinates {
		cell := l.geoIndex.Find(coordinate)
		candidates := l.geoPolygons[cell]
		props := make([]Props, 0, len(candidates))

		for _, candidate := range candidates {
			ll := s2.LatLngFromDegrees(coordinate.Latitude, coordinate.Longitude)
			point := s2.PointFromLatLng(ll)

			if candidate.Polygon.ContainsPoint(point) {
				props = append(props, candidate.Props)
			}
		}

		result = append(result, CoordinateProps{
			Props:      props,
			Coordinate: coordinate,
		})
	}

	return result
}
