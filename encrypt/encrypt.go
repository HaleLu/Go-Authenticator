package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"strings"
	"time"

	"github.com/HaleLu/go-authenticator/util"
)

func getCode(secret string, timestamp int64) (code int, err error) {
	var key, value []byte
	secret = strings.ToUpper(strings.Replace(secret, " ", "", -1))
	key, err = base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return
	}
	value = util.Int64ToBytes(timestamp / 30)

	hasher := hmac.New(sha1.New, key)
	hasher.Write(value)
	hash := hasher.Sum(nil)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	code, err = util.BytesToInt(hashParts)
	if err != nil {
		return
	}
	code = code % 1000000
	return
}

// GetCodeNow hash now timestamp with secret into a six-digit number as the auth code.
func GetCodeNow(secret string) (code int, err error) {
	return getCode(secret, time.Now().Unix())
}

// GetCode hash timestamp with secret into a six-digit number as the auth code.
func GetCode(secret string, timestamp int64) (code int, err error) {
	return getCode(secret, timestamp)
}

// CheckCode check whether code match the secret
func CheckCode(secret string, code int) (res bool, err error) {
	var _code int
	if _code, err = getCode(secret, time.Now().Unix()); err != nil {
		return
	}
	res = code == _code
	return
}
