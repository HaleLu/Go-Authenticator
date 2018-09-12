package qrcode

import (
	"fmt"
)

const format = "otpauth://totp/%s:%s?secret=%s&issuer=%s"

// GetURI gets a standard format key uri.
func GetURI(provider, user, secret string) string {
	return fmt.Sprintf(format, provider, user, secret, provider)
}
