package encrypt

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSignature(test *testing.T) {
	expect := "HCbG3xNE3vzhO%2Bu7qCUL1jS5hsu2n5r2cFhnTrtyDAE%3D"
	sign := Signature("1546084445901", "testappSecret")
	fmt.Println(sign)
	ioutil.WriteFile(".name", []byte("hello"), 0644)
	if sign != expect {
		test.Fatalf("期望值是:%s,结果值是:%s", expect, sign)
	}
}
