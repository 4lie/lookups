package lookups

import (
	"github.com/golang/geo/s2"
)

const (
	// Default level of s2 cells.
	DefaultS2CellLevel = 15
)

//nolint:gochecknoglobals
var (
	// Default value for geo indexer.
	DefaultGeoIndexer = NewS2Index(DefaultS2CellLevel)
)

type (
	// Lookuper is an interface for lookup services.
	Lookuper interface {
		Lookup(coordinates []Coordinate) []CoordinateProps
	}

	// lookups is an engine of lookups.
	lookups struct {
		geoIndex    GeoIndex
		geoPolygons map[string][]PolyProps
	}
)

// New returns a new lookups engine instance.
func New(polyProps []PolyProps, geoIndex GeoIndex) Lookuper {
	geoPolygons := make(map[string][]PolyProps)

	for _, polyProp := range polyProps {
		ids := geoIndex.Cover(polyProp.Polygon)

		for _, id := range ids {
			geoPolygons[id] = append(geoPolygons[id], polyProp)
		}
	}

	return &lookups{
		geoIndex:    geoIndex,
		geoPolygons: geoPolygons,
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
