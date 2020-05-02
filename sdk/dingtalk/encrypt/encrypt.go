package encrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
)

/**

钉钉的签名：https://ding-doc.dingtalk.com/doc#/faquestions/hxs5v9
签名算法为HmacSHA256，签名数据是当前时间戳timestamp，密钥是appId对应的appSecret

timestamp : 当前时间戳，单位是毫秒
appSecret : 扫码登录应用的appId

签名例子参考:
timestamp=1546084445901
appSecret=testappSecret
返回:HCbG3xNE3vzhO%2Bu7qCUL1jS5hsu2n5r2cFhnTrtyDAE%3D
*/
func Signature(timestamp, appSecret string) string {
	key := []byte(appSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(timestamp))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	// 特殊字符转换
	return url.QueryEscape(sign)
}
