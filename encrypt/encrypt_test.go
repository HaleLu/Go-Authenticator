package encrypt

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetCode(t *testing.T) {
	Convey("test GetCode", t, func() {
		now := time.Now().Unix()
		code, err := GetCode("0G6969V5CVN09JKQ", now)
		So(err, ShouldBeNil)
		So(code, ShouldBeGreaterThanOrEqualTo, 0)
	})
}
