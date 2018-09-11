package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Int64ToBytes(t *testing.T) {
	Convey("test Int64ToBytes", t, func() {
		i := int64(0x123456789abcdef0)
		bytes := Int64ToBytes(i)
		So(bytes[0], ShouldEqual, 0x12)
		So(bytes[1], ShouldEqual, 0x34)
		So(bytes[2], ShouldEqual, 0x56)
		So(bytes[3], ShouldEqual, 0x78)
		So(bytes[4], ShouldEqual, 0x9a)
		So(bytes[5], ShouldEqual, 0xbc)
		So(bytes[6], ShouldEqual, 0xde)
		So(bytes[7], ShouldEqual, 0xf0)
	})
}

func Test_BytesToUint32(t *testing.T) {
	Convey("test Int64ToBytes", t, func() {
		bytes := []byte{0x12, 0x34, 0x56, 0x78}
		i, err := BytesToUint32(bytes)
		So(err, ShouldBeNil)
		So(i, ShouldEqual, 0x12345678)
	})
}
