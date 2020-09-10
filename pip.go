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

// NewPipEngine create a new point in polygon engine using geom library.
func NewPipEngineFromGeom(polygons []GeomPolyProps) PipEngine {
	return PipEngine{}
}

// NewPipEngine create a new point in polygon engine using s2 library.
func NewPipEngineFromS2(polygons []S2PolyProps) PipEngine {
	return PipEngine{}
}

// Lookup return list of properties for the list of coordinates.
func (e PipEngine) Lookup(coordinates []Coordinate) []CoordinateProps {
	return []CoordinateProps{}
}
