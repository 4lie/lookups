package lookups

import (
	"sync"

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

	// Engine is an engine of lookups.
	Engine struct {
		s2Hash index.S2Hash
		mu     sync.RWMutex
		items  map[string][]PolyProps
	}
)

// New create a new lookups instance.
func New(s2Level int) *Engine {
	return &Engine{
		s2Hash: index.NewS2Hash(s2Level),
		items:  make(map[string][]PolyProps),
	}
}

// Set sets the given polygon and it's properties.
func (e *Engine) Set(polyProp PolyProps) {
	e.mu.Lock()
	defer e.mu.Unlock()

	ids := e.s2Hash.Cover(polyProp.Polygon)

	for _, id := range ids {
		e.items[id] = append(e.items[id], polyProp)
	}
}

// Lookup returns list of properties of given coordinates.
func (e *Engine) Lookup(coordinates []Coordinate) []CoordinateProps {
	e.mu.RLock()
	defer e.mu.RUnlock()

	out := make([]CoordinateProps, 0, len(coordinates))

	for _, coordinate := range coordinates {
		cell := e.cell(coordinate.Latitude, coordinate.Longitude)
		candidates := e.items[cell]
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

// cell find hex-encoded cell id of the given coordinate
// according to level specified in the New function.
func (e *Engine) cell(lat, lng float64) string {
	ll := s2.LatLngFromDegrees(lat, lng)
	level := e.s2Hash.Level()

	return s2.CellIDFromLatLng(ll).Parent(level).ToToken()
}
