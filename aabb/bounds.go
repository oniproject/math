package aabb

import "github.com/oniproject/math/f"

type Bounds struct {
	Min, Max f.Float
}

func BoundsOverlap(a, b Bounds) bool {
	return (a.Min <= b.Max && b.Min <= a.Max)
}

func ToBounds(bb AABB) Bounds {
	return Bounds{bb.L, bb.R}
}
