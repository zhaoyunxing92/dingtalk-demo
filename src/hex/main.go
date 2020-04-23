package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {

	sing := hmacSha256("1546084445901", "testappSecret")
	fmt.Println(sing == "HCbG3xNE3vzhO+u7qCUL1jS5hsu2n5r2cFhnTrtyDAE=")
	fmt.Println(sing)
	sing = strings.ReplaceAll(sing, "+", "%20")
	sing = strings.ReplaceAll(sing, "*", "%2A")
	sing = strings.ReplaceAll(sing, "~", "%7E")
	sing = strings.ReplaceAll(sing, "/", "%2F")
	sing = strings.ReplaceAll(sing, "=", "%3D")

	fmt.Println(sing)
}

func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hmacSha256(data, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
