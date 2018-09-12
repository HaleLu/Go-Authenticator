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

// BytesToInt converts []byte to int
func BytesToInt(bytes []byte) (value int, err error) {
	if len(bytes) != 4 {
		err = errors.New("length of bytes must be 4")
	}
	for i := 0; i < 4; i++ {
		value = value << 8
		value += int(bytes[i])
	}
	return
}

// randomBase32Char generates a random base32 character
func randomBase32Char() byte {
	r := _rand.Intn(32)
	if r < 26 {
		return 'A' + byte(r)
	}
	return '2' + byte(r) - 26
}

// RandomBase32String generates a random base32 string
func RandomBase32String() string {
	bs := make([]byte, 16)
	for i := range bs {
		bs[i] = randomBase32Char()
	}
	return string(bs)
}
