package qrcode

import (
	"encoding/base32"
	"errors"

	"github.com/HaleLu/go-authenticator/util"

	xqrcode "github.com/skip2/go-qrcode"
)

// GetQR generates a byte array which present an auth uri png.
func GetQR(provider, user string, key []byte) (png []byte, err error) {
	var (
		secret string
		uri    string
	)
	if len(key) != 20 {
		err = errors.New("length of key must be 20")
	}
	secret = base32.StdEncoding.EncodeToString(key)
	uri = GetURI(provider, user, secret)
	png, err = xqrcode.Encode(uri, xqrcode.Medium, 256)
	return
}

// RandomKey generates a random 160-bits key
func RandomKey() []byte {
	return util.RandomBytes(20)
}
