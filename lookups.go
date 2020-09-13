package lookups

import (
	"github.com/4lie/lookups/index"
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
	Lookups struct {
		index index.S2
		items map[string][]PolyProps
	}
)

// New create a new Lookups instance.
func New(polyProps []PolyProps, s2Level int) *Lookups {
	i := index.NewS2(s2Level)
	items := make(map[string][]PolyProps)

	for _, polyProp := range polyProps {
		ids := i.Cover(polyProp.Polygon)
		for _, id := range ids {
			items[id] = append(items[id], polyProp)
		}
	}

	return &Lookups{
		index: i,
		items: items,
	}
}

// Lookup returns list of properties of given coordinates.
func (l *Lookups) Lookup(coordinates []Coordinate) []CoordinateProps {
	out := make([]CoordinateProps, 0, len(coordinates))

	for _, coordinate := range coordinates {
		cell := l.cell(coordinate.Latitude, coordinate.Longitude)
		candidates := l.items[cell]
		props := make([]Props, 0, len(candidates))

		for _, candidate := range candidates {
			ll := s2.LatLngFromDegrees(coordinate.Latitude, coordinate.Longitude)
			point := s2.PointFromLatLng(ll)

			if candidate.Polygon.ContainsPoint(point) {
				props = append(props, candidate.Props)
			}
		}

		out = append(out, CoordinateProps{
			Props:      props,
			Coordinate: coordinate,
		})
	}

	return out
}

// cell finds hex-encoded cell id of the given coordinate according to level specified in the New function.
func (l *Lookups) cell(lat, lng float64) string {
	ll := s2.LatLngFromDegrees(lat, lng)
	level := l.index.Level()

	return s2.CellIDFromLatLng(ll).Parent(level).ToToken()
}
