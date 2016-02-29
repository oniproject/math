package t

import "github.com/oniproject/math/f"
import "github.com/oniproject/math/v"
import "github.com/oniproject/math/aabb"

type Transform struct {
	A, B, C, D, Tx, Ty f.Float
}

// Transform an absolute point. (i.e. a vertex)
func (t *Transform) Point(p v.Vect) v.Vect {
	return v.Vect{
		t.A*p.X + t.C*p.Y + t.Tx,
		t.B*p.X + t.D*p.Y + t.Ty,
	}
}
func (t *Transform) PointInverse(p v.Vect) v.Vect {
	id := 1.0 / (t.A*t.D - t.C*t.B)
	x, y := p.X, p.Y
	return v.Vect{
		t.D*id*x + -t.C*id*y + (+t.Ty*t.C-t.Tx*t.D)*id,
		t.A*id*y + -t.B*id*x + (-t.Ty*t.A+t.Tx*t.B)*id,
	}
}

// Transform a vector (i.e. a normal)
func (t *Transform) Vect(p v.Vect) v.Vect {
	return v.Vect{
		t.A*p.X + t.C*p.Y,
		t.B*p.X + t.D*p.Y,
	}
}

// Identity transform matrix.
func Identity() Transform {
	return Transform{1.0, 0.0, 0.0, 1.0, 0.0, 0.0}
}

// Construct a new transform matrix.
// (a, b) is the x basis vector.
// (c, d) is the y basis vector.
// (tx, ty) is the translation.
func New(a, b, c, d, tx, ty f.Float) Transform {
	return Transform{a, b, c, d, tx, ty}
}

// Construct a new transform matrix in transposed order.
func Transpose(a, c, tx, b, d, ty f.Float) Transform {
	return Transform{a, b, c, d, tx, ty}
}

// Get the inverse of a transform matrix.
func Inverse(t Transform) Transform {
	inv_det := 1.0 / (t.A*t.D - t.C*t.B)
	return Transpose(
		+t.D*inv_det, -t.C*inv_det, (t.C*t.Ty-t.Tx*t.D)*inv_det,
		-t.B*inv_det, +t.A*inv_det, (t.Tx*t.B-t.A*t.Ty)*inv_det,
	)
}

// Multiply two transformation matrices.
func Mult(t1, t2 Transform) Transform {
	return Transpose(
		t1.A*t2.A+t1.C*t2.B, t1.A*t2.C+t1.C*t2.D, t1.A*t2.Tx+t1.C*t2.Ty+t1.Tx,
		t1.B*t2.A+t1.D*t2.B, t1.B*t2.C+t1.D*t2.D, t1.B*t2.Tx+t1.D*t2.Ty+t1.Ty,
	)
}

// Transform a cpBB.
func (t *Transform) BB(bb aabb.AABB) aabb.AABB {
	center := bb.Center()
	hw := (bb.R - bb.L) * 0.5
	hh := (bb.T - bb.B) * 0.5

	a, b, d, e := t.A*hw, t.C*hh, t.B*hw, t.D*hh
	hw_max := f.Max(f.Abs(a+b), f.Abs(a-b))
	hh_max := f.Max(f.Abs(d+e), f.Abs(d-e))
	return aabb.ForExtents(t.Point(center), hw_max, hh_max)
}

// Create a transation matrix.
func Translate(translate v.Vect) Transform {
	return Transpose(
		1.0, 0.0, translate.X,
		0.0, 1.0, translate.Y,
	)
}

// Create a scale matrix.
func Scale(scaleX, scaleY f.Float) Transform {
	return Transpose(
		scaleX, 0.0, 0.0,
		0.0, scaleY, 0.0,
	)
}

// Create a rotation matrix.
func Rotate(radians f.Float) Transform {
	rot := v.ForAngle(radians)
	return Transpose(
		rot.X, -rot.Y, 0.0,
		rot.Y, +rot.X, 0.0,
	)
}

// Create a rigid transformation matrix. (transation + rotation)
func Rigid(translate v.Vect, radians f.Float) Transform {
	rot := v.ForAngle(radians)
	return Transpose(
		rot.X, -rot.Y, translate.X,
		rot.Y, +rot.X, translate.Y,
	)
}

// Fast inverse of a rigid transformation matrix.
func RigidInverse(t Transform) Transform {
	return Transpose(
		+t.D, -t.C, (t.C*t.Ty - t.Tx*t.D),
		-t.B, +t.A, (t.Tx*t.B - t.A*t.Ty),
	)
}

// XXX: Miscellaneous (but useful) transformation matrices.
// See source for documentation...

func Wrap(outer, inner Transform) Transform {
	return Mult(Inverse(outer), Mult(inner, outer))
}

func WrapInverse(outer, inner Transform) Transform {
	return Mult(outer, Mult(inner, Inverse(outer)))
}

func Ortho(bb aabb.AABB) Transform {
	return Transpose(
		2.0/(bb.R-bb.L), 0.0, -(bb.R+bb.L)/(bb.R-bb.L),
		0.0, 2.0/(bb.T-bb.B), -(bb.T+bb.B)/(bb.T-bb.B),
	)
}

func BoneScale(v0, v1 v.Vect) Transform {
	d := v.Sub(v1, v0)
	return Transpose(
		d.X, -d.Y, v0.X,
		d.Y, +d.X, v0.Y,
	)
}

func AxialScale(axis, pivot v.Vect, scale f.Float) Transform {
	A := axis.X * axis.Y * (scale - 1.0)
	B := v.Dot(axis, pivot) * (1.0 - scale)

	return Transpose(
		scale*axis.X*axis.X+axis.Y*axis.Y, A, axis.X*B,
		A, axis.X*axis.X+scale*axis.Y*axis.Y, axis.Y*B,
	)
}
