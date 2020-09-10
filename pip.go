package lookups

type (
	// Lookuper Is an interface for lookup services.
	Lookuper interface {
		Lookup(coordinates []Coordinate) []CoordinateProps
	}

	// PipEngine Is an engine for point in polygon.
	PipEngine struct {
	}
)

// NewPipEngine Create a new point in polygon engine.
func NewPipEngine(polygons []PolyProps) PipEngine {
	return PipEngine{}
}

// Lookup Return list of properties for the list of coordinates.
func (e PipEngine) Lookup(coordinates []Coordinate) []CoordinateProps {
	return []CoordinateProps{}
}
