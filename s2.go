package lookups

import (
	"github.com/golang/geo/s2"
)

type S2Hash struct {
	level int
}

func NewS2Hash(level int) S2Hash {
	return S2Hash{level: level}
}

func (s S2Hash) Cover(polygon *s2.Polygon) []string {
	if polygon.NumEdges() < 1 {
		return []string{}
	}

	ids := s2.SimpleRegionCovering(polygon, polygon.Edge(0).V0, s.level)

	out := make([]string, 0, len(ids))

	for _, id := range ids {
		out = append(out, id.ToToken())
	}

	return out
}
