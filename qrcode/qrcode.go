package qrcode

import (
	"errors"

	"github.com/HaleLu/go-authenticator/util"

	xqrcode "github.com/skip2/go-qrcode"
)

// GetQR generates a byte array which present an auth uri png.
func GetQR(provider, user, secret string) (png []byte, err error) {
	if len(secret) != 16 {
		err = errors.New("length of secret must be 20")
	}
	png, err = xqrcode.Encode(GetURI(provider, user, secret), xqrcode.Medium, 256)
	return
}

// RandomKey generates a random 160-bits key
func RandomKey() string {
	return util.RandomBase32String()
}
