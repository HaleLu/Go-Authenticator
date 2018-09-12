package qrcode

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetQR(t *testing.T) {
	Convey("test GetQR", t, func() {
		png, err := GetQR("Bilibili", "luhao", RandomKey())
		So(err, ShouldBeNil)
		err = ioutil.WriteFile("./qr.png", png, 0777) //写入文件(字节数组)
		So(err, ShouldBeNil)
	})
}
