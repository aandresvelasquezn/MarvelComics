package tools

import (
	"crypto/md5"
	"fmt"
)

func digestString(t string, publicKey string, privateKey string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(t+privateKey+publicKey)))
}

// GenerateHash Generate Hash code
func GenerateHash() string {
	t := "1"
	privateKey := GetDotEnvVariable("PRIVATE_KEY")
	publicKey := GetDotEnvVariable("PUBLIC_KEY")
	return digestString(t, publicKey, privateKey)
}
