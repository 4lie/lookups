package lookups

const (
	// Default level of s2 cells.
	DefaultS2CoverLevel = 15
)

type (
	// Lookuper is an interface for lookup services.
	Lookuper interface {
		Lookup(coordinates []Coordinate) []CoordinateProps
	}

	// PipEngine is an engine for point in polygon.
	PipEngine struct {
		S2Hash S2Hash
	}
)

// NewPipEngine create a new point in polygon engine using.
func NewPipEngine(polygons []PolyProps) PipEngine {
	return PipEngine{
		S2Hash: NewS2Hash(DefaultS2CoverLevel),
	}
}

// Lookup return list of properties for the list of coordinates.
func (e PipEngine) Lookup(coordinates []Coordinate) []CoordinateProps {
	return []CoordinateProps{}
}
