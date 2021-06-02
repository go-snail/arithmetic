package leecode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWaterKing(t *testing.T) {
	convey.Convey("TestWaterKing",t, func() {
		src := []int{-1,-1,3}
		ret := waterKing(src)
		convey.So(ret,convey.ShouldEqual,3)
	})
}
