package f

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func Test(test *testing.T) {
	Convey("Floating", test, func() {
		Convey("From math", func() {
			So(Sqrt(9), ShouldEqual, math.Sqrt(9))
			So(Sin(9), ShouldEqual, math.Sin(9))
			So(Cos(9), ShouldEqual, math.Cos(9))
			So(Acos(1), ShouldEqual, math.Acos(1))
			So(Atan2(9, 5), ShouldEqual, math.Atan2(9, 5))
			So(Mod(9, 5), ShouldEqual, math.Mod(9, 5))
			So(Exp(9), ShouldEqual, math.Exp(9))
			So(Pow(9, 5), ShouldEqual, math.Pow(9, 5))
			So(Floor(9.4), ShouldEqual, math.Floor(9.4))
			So(Ceil(9.4), ShouldEqual, math.Ceil(9.4))

			So(IsNaN(Float(math.NaN())), ShouldBeTrue)
			So(IsNaN(0), ShouldBeFalse)
			So(IsNaN(1), ShouldBeFalse)
		})

		Convey("Max", func() {
			var a, b Float = 1, 99
			So(Max(a, b), ShouldEqual, b)
		})
		Convey("Min", func() {
			var a, b Float = 1, 99
			So(Min(a, b), ShouldEqual, a)
		})
		Convey("Abs", func() {
			So(Abs(0), ShouldEqual, 0)
			So(Abs(+1), ShouldEqual, 1)
			So(Abs(-1), ShouldEqual, 1)
		})
		Convey("Clamp", func() {
			var min, max Float = 0.0, 1.0
			So(Clamp(+0.5, min, max), ShouldEqual, 0.5)
			So(Clamp(-5.0, min, max), ShouldEqual, 0.0)
			So(Clamp(+5.0, min, max), ShouldEqual, 1.0)
		})
		Convey("Clamp01", func() {
			So(Clamp01(+0.5), ShouldEqual, 0.5)
			So(Clamp01(-5.0), ShouldEqual, 0.0)
			So(Clamp01(+5.0), ShouldEqual, 1.0)
		})

		Convey("Lerp", func() {
			var min, max Float = 0.0, 2.0
			So(Lerp(min, max, +0), ShouldEqual, 0.0)
			So(Lerp(min, max, +1), ShouldEqual, 2.0)
			So(Lerp(min, max, +2), ShouldEqual, 4.0)
			So(Lerp(min, max, -1), ShouldEqual, -2.0)
		})
		Convey("LerpConst", func() {
			// TODO
			var min, max Float = 0.0, 2.0
			So(LerpConst(min, max, +0), ShouldEqual, 0.0)
			So(LerpConst(min, max, +1), ShouldEqual, 1.0)
			So(LerpConst(min, max, +2), ShouldEqual, 2.0)
			So(LerpConst(min, max, -1), ShouldEqual, -1.0)
			So(LerpConst(min, max, -2), ShouldEqual, -2.0)
			So(LerpConst(min, max, -4), ShouldEqual, -4.0)
		})

	})
}
