package lookups

type (
	Lookuper interface {
		Lookup(coordinates []Coordinate) ([]CoordinateProps, error)
	}

	PipEngine struct {
	}
)

// NewPipEngine Create a new point in polygon engine.
func NewPipEngine(polygons []PolyProps) (*PipEngine, error) {
	return nil, nil
}

func (e *PipEngine) Lookup(coordinates []Coordinate) ([]CoordinateProps, error) {
	return nil, nil
}
