package dingtalk

import (
	"fmt"
	"testing"
)

func TestSignature(t *testing.T) {
	sign := Signature("1546084445901", "testappSecret")
	fmt.Println(sign)
	if sign != "HCbG3xNE3vzhO%2Bu7qCUL1jS5hsu2n5r2cFhnTrtyDAE%3D" {
		panic("钉钉签名异常")
	}
}

func TestGetUserInfoByCode(t *testing.T) {
	GetUserInfoByCode("5555")
}

func TestGetAccessToken(t *testing.T) {
	token, ex := GetAccessToken()
	fmt.Println(token, ex)
}
