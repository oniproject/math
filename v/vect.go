package v

import "github.com/oniproject/math/f"

// Chipmunk's 2D vector type along with a handy 2D vector math lib.
type Vect struct{ X, Y f.Float }

// Constant for the zero vector.
func Zero() Vect { return Vect{} }

// Convenience constructor for Vect structs.
func V(x, y f.Float) Vect { return Vect{x, y} }

// Check if two vectors are equal.
// (Be careful when comparing floating point numbers!)
func Eql(v1, v2 Vect) bool { return v1.X == v2.X && v1.Y == v2.Y }

// Add two vectors
func Add(v1, v2 Vect) Vect { return Vect{v1.X + v2.X, v1.Y + v2.Y} }

// Subtract two vectors.
func Sub(v1, v2 Vect) Vect { return Vect{v1.X - v2.X, v1.Y - v2.Y} }

// Negate a vector.
func Neg(v Vect) Vect { return Vect{-v.X, -v.Y} }

// Scalar multiplication.
func Mult(v Vect, s f.Float) Vect { return Vect{v.X * s, v.Y * s} }

// Vector dot product.
func Dot(v1, v2 Vect) f.Float { return v1.X*v2.X + v1.Y*v2.Y }

// 2D vector cross product analog.
// The cross product of 2D vectors results in a 3D vector with only a z component.
// This function returns the magnitude of the z value.
func Cross(v1, v2 Vect) f.Float { return v1.X*v2.Y - v1.Y*v2.X }

// Returns a perpendicular vector. (90 degree rotation)
//func Perp(v Vect) Vect { return Vect{-v.Y, v.X} }

// Returns a perpendicular vector. (90 degree rotation)
func LPerp(v Vect) Vect { return Vect{-v.Y, v.X} }

// Returns a perpendicular vector. (-90 degree rotation)
func RPerp(v Vect) Vect { return Vect{v.Y, -v.X} }

// Returns the vector projection of v1 onto v2.
func Project(v1, v2 Vect) Vect {
	return Mult(v2, Dot(v1, v2)/Dot(v2, v2))
}

// Returns the unit length vector for the given angle (in radians).
func ForAngle(a f.Float) Vect { return Vect{f.Cos(a), f.Sin(a)} }

// Returns the angular direction v is pointing in (in radians).
func ToAngle(v Vect) f.Float { return f.Atan2(v.Y, v.X) }

// Uses complex number multiplication to rotate v1 by v2.
// Scaling will occur if v1 is not a unit vector.
func Rotate(v1, v2 Vect) Vect {
	return Vect{v1.X*v2.X - v1.Y*v2.Y, v1.X*v2.Y + v1.Y*v2.X}
}

// Inverse of Rotate().
func UnRotate(v1, v2 Vect) Vect {
	return Vect{v1.X*v2.X + v1.Y*v2.Y, v1.Y*v2.X - v1.X*v2.Y}
}

// Returns the squared length of v.
// Faster than Length() when you only need to compare lengths.
func LengthSq(v Vect) f.Float { return Dot(v, v) }

// Returns the length of v.
func Length(v Vect) f.Float { return f.Sqrt(Dot(v, v)) }

// Linearly interpolate between v1 and v2.
func Lerp(v1, v2 Vect, t f.Float) Vect {
	return Add(Mult(v1, 1.0-t), Mult(v2, t))
}

// Returns a normalized copy of v.
func Normalize(v Vect) Vect {
	// Neat trick I saw somewhere to avoid div/0.
	return Mult(v, 1.0/(Length(v)+f.FloatMin))
}

// Spherical linearly interpolate between v1 and v2.
func Slerp(v1, v2 Vect, t f.Float) Vect {
	dot := Dot(Normalize(v1), Normalize(v2))
	omega := f.Acos(f.Clamp(dot, -1.0, 1.0))

	if omega < 1e-3 {
		// If the angle between two vectors is very small, lerp instead to avoid precision issues.
		return Lerp(v1, v2, t)
	} else {
		denom := 1.0 / f.Sin(omega)
		return Add(
			Mult(v1, f.Sin((1.0-t)*omega)*denom),
			Mult(v2, f.Sin(t*omega)*denom),
		)
	}
}

// Spherical linearly interpolate between v1 towards v2 by no more than angle a radians
func SlerpConst(v1, v2 Vect, a f.Float) Vect {
	dot := Dot(Normalize(v1), Normalize(v2))
	omega := f.Acos(f.Clamp(dot, -1.0, 1.0))

	return Slerp(v1, v2, f.Min(a, omega)/omega)
}

// Clamp v to length l.
func Clamp(v Vect, l f.Float) Vect {
	if Dot(v, v) > l*l {
		return Mult(Normalize(v), l)
	}
	return v
}

// Linearly interpolate between v1 towards v2 by distance d.
func LerpConst(v1, v2 Vect, d f.Float) Vect {
	return Add(v1, Clamp(Sub(v2, v1), d))
}

// Returns the distance between v1 and v2.
func Dist(v1, v2 Vect) f.Float { return Length(Sub(v1, v2)) }

// Returns the squared distance between v1 and v2. Faster than Dist() when you only need to compare distances.
func DistSq(v1, v2 Vect) f.Float { return LengthSq(Sub(v1, v2)) }

// Returns true if the distance between v1 and v2 is less than dist.
func Near(v1, v2 Vect, dist f.Float) bool {
	return DistSq(v1, v2) < dist*dist
}

// Check if two vectors are equal.
// (Be careful when comparing floating point numbers!)
func (p *Vect) Eql(q Vect) bool {
	return p.X == p.X && p.Y == p.Y
}

// Add two vectors
func (p *Vect) Add(q Vect) {
	p.X += q.X
	p.Y += q.Y
}

// Subtract two vectors.
func (p *Vect) Sub(q Vect) {
	p.X -= q.X
	p.Y -= q.Y
}

// Negate a vector.
func (p *Vect) Neg() {
	p.X = -p.X
	p.Y = -p.Y
}

// Scalar multiplication.
func (p *Vect) Mult(s f.Float) {
	p.X *= s
	p.Y *= s
}

// Returns the angular direction v is pointing in (in radians).
func (p *Vect) ToAngle() f.Float { return f.Atan2(p.Y, p.X) }

// Returns the squared length of v.
// Faster than Length() when you only need to compare lengths.
func (p Vect) LengthSq() f.Float { return Dot(p, p) }

// Returns the length of v.
func (p Vect) Length() f.Float { return f.Sqrt(Dot(p, p)) }

// Returns a normalized copy of v.
func (p *Vect) Normalize() {
	// Neat trick I saw somewhere to avoid div/0.
	p.Mult(1.0 / (p.Length() + f.FloatMin))
}

// Clamp v to length l.
func (p *Vect) Clamp(l f.Float) {
	if Dot(*p, *p) > l*l {
		p.Normalize()
		p.Mult(l)
	}
}
