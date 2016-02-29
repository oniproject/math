package t

import (
	"github.com/oniproject/math/v"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAABB(test *testing.T) {
	Convey("Transform", test, func() {
		t, s, r := Translate(v.Vect{2, 3}), Scale(4, 5), Rotate(0.5)

		Convey("Inverse", func() {
			ti, si, ri := Inverse(t), Inverse(s), Inverse(r)
			Convey("Point", func() {
				var p v.Vect
				p = v.V(-1, 8)

				So(t.PointInverse(p), ShouldResemble, ti.Point(p))
				So(s.PointInverse(p), ShouldResemble, si.Point(p))
				So(r.PointInverse(p), ShouldResemble, ri.Point(p))
			})

			SkipConvey("Vect", func() {
				//var p v.Vect
				//p = v.V(-1, 0)

				//So(t.VectInverse(p), ShouldResemble, ti.Vect(p))
				// FIXME So(s.VectInverse(p), ShouldResemble, si.Vect(p))
				// FIXME So(r.VectInverse(p), ShouldResemble, ri.Vect(p))
			})
		})
	})
}
