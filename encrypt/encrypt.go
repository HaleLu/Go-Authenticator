package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"

	"github.com/HaleLu/go-authenticator/util"
)

// GetCode hash value with key into a six-digit number as the auth code.
func GetCode(key []byte, value []byte) (code uint32) {
	hasher := hmac.New(sha1.New, key)
	hasher.Write(value)
	hash := hasher.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number, _ := util.BytesToUint32(hashParts)
	code = number % 1000000
	return
}
