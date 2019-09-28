package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// HMACSHA256Base64Sign sign with hmac sha256, base64 encode
func HMACSHA256Base64Sign(secret, params string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	signByte := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signByte), nil
}
