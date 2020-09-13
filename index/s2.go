package index

import (
	"github.com/golang/geo/s2"
)

// S2 is a S2-based polygon indexer.
type S2 struct {
	level int
}

// NewS2 returns a new S2 instance.
func NewS2(level int) S2 {
	return S2{level: level}
}

// Covers returns all S2 cells inside the given polygon.
func (s S2) Cover(polygon *s2.Polygon) []string {
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

// Level returns S2 cell level.
func (s S2) Level() int {
	return s.level
}
