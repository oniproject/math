package v

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

type addTest struct {
	in1, in2 Vect
	out      Vect
}

/*
type minTest struct {
	in1, in2 Vect
	out      Vect
}
type maxTest struct {
	in1, in2 Vect
	out      Vect
}
*/
type distTest struct {
	in1, in2 Vect
	out      float64
}

func Test(test *testing.T) {
	Convey("Vector", test, func() {
		Convey("Zero", func() {
			So(Zero(), ShouldResemble, Vect{0, 0})
		})
		Convey("V", func() {
			So(V(1, 2), ShouldResemble, Vect{1, 2})
		})
		Convey("Neg", func() {
			So(Neg(Vect{+1, +2}), ShouldResemble, Vect{-1, -2})
			So(Neg(Vect{-1, -2}), ShouldResemble, Vect{+1, +2})
			So(Neg(Vect{+1, -2}), ShouldResemble, Vect{-1, +2})
			So(Neg(Vect{-1, +2}), ShouldResemble, Vect{+1, -2})
		})
		Convey("Add", func() {
			for _, at := range addTests {
				So(Add(at.in1, at.in2), ShouldResemble, at.out)
				So(Eql(Add(at.in1, at.in2), at.out), ShouldBeTrue)
			}
		})
		Convey("Dist", func() {
			for _, at := range distTests {
				So(Dist(at.in1, at.in2), ShouldEqual, at.out)
			}
		})

		Convey("v.Clamp", func() {
			v1 := Vect{5, 0}
			v1.Clamp(2)
			So(v1, ShouldResemble, Vect{2, 0})

			v2 := Vect{0, 5}
			v2.Clamp(2)
			So(v2, ShouldResemble, Vect{0, 2})
			Convey("v.Length", func() {
				So(v1.Length(), ShouldEqual, 2)
				So(v2.Length(), ShouldEqual, 2)
			})
			Convey("v.LengthSq", func() {
				So(v1.LengthSq(), ShouldEqual, 2*2)
				So(v2.LengthSq(), ShouldEqual, 2*2)
			})
		})
	})
}

var addTests = []addTest{
	{Vect{0, 0}, Vect{0, 0}, Vect{0, 0}},
	{Vect{0, 1}, Vect{0, 0}, Vect{0, 1}},
	{Vect{1, 0}, Vect{0, 0}, Vect{1, 0}},
	{Vect{1, 2}, Vect{0, 0}, Vect{1, 2}},
	{Vect{0, 0}, Vect{0, 1}, Vect{0, 1}},
	{Vect{0, 0}, Vect{1, 0}, Vect{1, 0}},
	{Vect{0, 0}, Vect{1, 2}, Vect{1, 2}},
	{Vect{2, 4}, Vect{1, 3}, Vect{3, 7}},
	{Vect{3, 1}, Vect{4, 2}, Vect{7, 3}},
	{Vect{2, 4}, Vect{2, 4}, Vect{4, 8}},
	{Vect{5, 5}, Vect{2, 2}, Vect{7, 7}},
}

/*
var minTests = []minTest{
	{Vect{0, 0}, Vect{0, 0}, Vect{0, 0}},
	{Vect{1, 2}, Vect{9, 9}, Vect{1, 2}},
	{Vect{9, 9}, Vect{1, 2}, Vect{1, 2}},
	{Vect{5, 2}, Vect{1, 4}, Vect{1, 2}},
	{Vect{9, 6}, Vect{7, 8}, Vect{7, 6}},
}
var maxTests = []maxTest{
	{Vect{0, 0}, Vect{0, 0}, Vect{0, 0}},
	{Vect{1, 2}, Vect{9, 9}, Vect{9, 9}},
	{Vect{9, 9}, Vect{1, 2}, Vect{9, 9}},
	{Vect{5, 2}, Vect{1, 4}, Vect{5, 4}},
	{Vect{9, 6}, Vect{7, 8}, Vect{9, 8}},
}
*/
var distTests = []distTest{
	{Vect{0, 0}, Vect{0, 0}, 0},
	{Vect{0, 2}, Vect{0, 0}, 2},
	{Vect{2, 0}, Vect{0, 0}, 2},
	{Vect{0, 0}, Vect{4, 0}, 4},
	{Vect{0, 0}, Vect{0, 4}, 4},
	{Vect{1, 1}, Vect{0, 0}, math.Sqrt(2)},
	{Vect{1, 1}, Vect{2, 2}, math.Sqrt(2)},
}
