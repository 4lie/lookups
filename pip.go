package lookups

type (
	// Lookuper is an interface for lookup services.
	Lookuper interface {
		Lookup(coordinates []Coordinate) []CoordinateProps
	}

	// PipEngine is an engine for point in polygon.
	PipEngine struct {
	}
)

// NewPipEngine create a new point in polygon engine using.
func NewPipEngine(polygons []PolyProps) PipEngine {
	return PipEngine{}
}

// Lookup return list of properties for the list of coordinates.
func (e PipEngine) Lookup(coordinates []Coordinate) []CoordinateProps {
	return []CoordinateProps{}
}
