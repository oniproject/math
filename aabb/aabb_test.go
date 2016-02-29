package aabb

import (
	"github.com/oniproject/math/v"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAABB(test *testing.T) {
	Convey("AABB", test, func() {
		Convey("ForExtents", func() {
			So(ForExtents(v.V(4, 4), 2, 1), ShouldResemble, AABB{
				L: 2, B: 3,
				R: 6, T: 5,
			})
		})
		Convey("ForCircle", func() {
			So(ForCircle(v.V(4, 4), 2), ShouldResemble, AABB{
				L: 2, B: 2,
				R: 6, T: 6,
			})
		})
		Convey("Intersects", func() {
			So(Intersects(
				New(0, 0, 6, 6), New(0, 0, 6, 6)),
				ShouldBeTrue)
			So(Intersects(
				New(1, 1, 6, 6), New(-1, -1, 3, 3)),
				ShouldBeTrue)
			So(Intersects(
				New(4, 5, 6, 6), New(-1, -1, 3, 3)),
				ShouldBeFalse)
		})
		Convey("Contains", func() {
			bb := New(0, 0, 6, 6)
			So(bb.Contains(New(+1, +1, 3, 4)), ShouldBeTrue)
			So(bb.Contains(New(-1, -1, 3, 4)), ShouldBeFalse)
		})
		Convey("ContainsVect", func() {
			bb := New(0, 0, 6, 6)
			So(bb.ContainsVect(v.V(1, 1)), ShouldBeTrue)
			So(bb.ContainsVect(v.V(-1, 1)), ShouldBeFalse)
			So(bb.ContainsVect(v.V(1, -1)), ShouldBeFalse)
			So(bb.ContainsVect(v.V(7, 1)), ShouldBeFalse)
			So(bb.ContainsVect(v.V(1, 7)), ShouldBeFalse)
		})
		Convey("Merge", func() {
			bb := Merge(New(0, 1, 8, 9), New(-1, -2, 3, 4))
			So(bb, ShouldResemble, New(-1, -2, 8, 9))
			// TODO for all four
		})

		SkipConvey("Expand", func() {
			//bb := Expand(New(0, 1, 2, 3), New(-1, -2, 3, 4))
			//So(bb, ShouldResemble, New(-1, -2, 8, 9))
			// TODO for all four
		})
	})
}

/*
type AABB struct {
	L, B, R, T f.Float
}

// Convenience constructor for AABB structs.
func New(l, b, r, t f.Float) AABB {
	return AABB{l, b, r, t}
}


// Returns a bounding box that holds both @c bb and @c v.
func Expand(bb AABB, v v.Vect) AABB {
	return New(
		f.Min(bb.L, v.X),
		f.Min(bb.B, v.Y),
		f.Max(bb.R, v.X),
		f.Max(bb.T, v.Y),
	)
}

// Returns the center of a bounding box.
func (bb *AABB) Center() v.Vect {
	return v.Lerp(v.Vect{bb.L, bb.B}, v.Vect{bb.R, bb.T}, 0.5)
}

// Returns the area of the bounding box.
func (bb *AABB) Area() f.Float {
	return (bb.R - bb.L) * (bb.T - bb.B)
}

// Merges @c a and @c b and returns the area of the merged bounding box.
func MergedArea(a, b AABB) f.Float {
	rl := f.Max(a.R, b.R) - f.Min(a.L, b.L)
	tb := f.Max(a.T, b.T) - f.Min(a.B, b.B)
	return rl * tb
}

// Returns the fraction along the segment query the AABB is hit. Returns INFINITY if it doesn't hit.
func (bb *AABB) SegmentQuery(a, b v.Vect) f.Float {
	delta := v.Sub(b, a)
	tmin, tmax := -f.Inf, f.Inf

	if delta.X != 0.0 {
		t1 := (bb.L - a.X) / delta.X
		t2 := (bb.R - a.X) / delta.X
		tmin = f.Max(tmin, f.Min(t1, t2))
		tmax = f.Min(tmax, f.Max(t1, t2))
	}

	if delta.Y != 0.0 {
		t1 := (bb.B - a.Y) / delta.Y
		t2 := (bb.T - a.Y) / delta.Y
		tmin = f.Max(tmin, f.Min(t1, t2))
		tmax = f.Min(tmax, f.Max(t1, t2))
	}

	if tmin <= tmax && 0.0 <= tmax && tmin <= 1.0 {
		return f.Max(tmin, 0.0)
	}

	return f.Inf
}

// Return true if the bounding box intersects the line segment with ends @c a and @c b.
func (bb *AABB) IntersectsSegment(a, b v.Vect) bool {
	return bb.SegmentQuery(a, b) != f.Inf
}

// Clamp a vector to a bounding box.
func (bb *AABB) ClampVect(p v.Vect) v.Vect {
	return v.Vect{f.Clamp(p.X, bb.L, bb.R), f.Clamp(p.Y, bb.B, bb.T)}
}

// Wrap a vector to a bounding box.
func (bb *AABB) WrapVect(p v.Vect) v.Vect {
	dx := f.Abs(bb.R - bb.L)
	x := f.Mod(p.X-bb.L, dx)
	if x <= 0.0 {
		x += dx
	}

	dy := f.Abs(bb.T - bb.B)
	y := f.Mod(p.Y-bb.B, dy)
	if y <= 0.0 {
		y += dy
	}

	return v.Vect{x + bb.L, y + bb.B}
}

// Returns a bounding box offseted by @c v.
func Offset(bb AABB, p v.Vect) AABB {
	return New(
		bb.L+p.X,
		bb.B+p.Y,
		bb.R+p.X,
		bb.T+p.Y,
	)
}
*/
