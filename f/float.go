package f

import "math"

const (
	Pi       Float = math.Pi
	FloatMin Float = math.SmallestNonzeroFloat64
)

var Inf = Float(math.Inf(+1))

type Float float32

func Sqrt(x Float) Float     { return Float(math.Sqrt(float64(x))) }
func Sin(x Float) Float      { return Float(math.Sin(float64(x))) }
func Cos(x Float) Float      { return Float(math.Cos(float64(x))) }
func Acos(x Float) Float     { return Float(math.Acos(float64(x))) }
func Atan2(x, y Float) Float { return Float(math.Atan2(float64(x), float64(y))) }
func Mod(x, y Float) Float   { return Float(math.Mod(float64(x), float64(y))) }
func Exp(x Float) Float      { return Float(math.Exp(float64(x))) }
func Pow(x, y Float) Float   { return Float(math.Pow(float64(x), float64(y))) }
func Floor(x Float) Float    { return Float(math.Floor(float64(x))) }
func Ceil(x Float) Float     { return Float(math.Ceil(float64(x))) }
func IsNaN(x Float) bool     { return math.IsNaN(float64(x)) }

// Return the max of two Floats.
func Max(a, b Float) Float {
	if a > b {
		return a
	}
	return b
}

// Return the min of two Floats.
func Min(a, b Float) Float {
	if a < b {
		return a
	}
	return b
}

// Return the absolute value of a Float.
func Abs(f Float) Float {
	if f < 0 {
		return -f
	}
	return f
}

// Clamp f to be between min and max.
func Clamp(f, min, max Float) Float {
	return Min(Max(f, min), max)
}

// Clamp f to be between 0 and 1.
func Clamp01(f Float) Float {
	return Max(0.0, Min(f, 1.0))
}

// Linearly interpolate (or extrapolate) between f1 and f2 by t percent.
func Lerp(f1, f2, t Float) Float {
	return f1*(1.0-t) + f2*t
}

// Linearly interpolate from f1 to f2 by no more than d.
func LerpConst(f1, f2, d Float) Float {
	return f1 + Clamp(f2-f1, -d, d)
}
