package qrcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetQR(t *testing.T) {
	Convey("test GetQR", t, func() {
		_, err := GetQR("Bilibili", "luhao", RandomKey())
		So(err, ShouldBeNil)
	})
}
