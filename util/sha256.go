package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Sha256(str string, key string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(str))
	s := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(s)
}
