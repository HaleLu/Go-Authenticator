package util

import (
	"errors"
	"math/rand"
	"time"
)

var _rand *rand.Rand

func init() {
	_rand = rand.New(rand.NewSource(time.Now().Unix()))
}

// Int64ToBytes converts int64 to []byte
func Int64ToBytes(value int64) (bytes []byte) {
	bytes = make([]byte, 8)
	for i := 0; i < 8; i++ {
		bytes[7-i] = byte(value & 0xFF)
		value = value >> 8
	}
	return
}

// BytesToUint32 converts []byte to uint32
func BytesToUint32(bytes []byte) (value uint32, err error) {
	if len(bytes) != 4 {
		err = errors.New("length of bytes must be 4")
	}
	for i := 0; i < 4; i++ {
		value = value << 8
		value += uint32(bytes[i])
	}
	return
}

// RandomChar generates a random character
func RandomChar() byte {
	return byte(_rand.Intn(256))
}

// RandomBytes generates a random byte array
func RandomBytes(length int) []byte {
	bs := make([]byte, length)
	for i := range bs {
		bs[i] = RandomChar()
	}
	return bs
}
