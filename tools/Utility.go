package tools

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

func GetToday(format string) (todayString string) {
	today := time.Now()
	todayString = today.Format(format)
	return
}

func EncryptionSha256(data []byte) string {
	hash   	 :=  	sha256.New()
	hash.Write(data)
	Encrypt  := 	base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return Encrypt
}