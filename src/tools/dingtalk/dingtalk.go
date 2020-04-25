package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/zhaoyunxing90/dingtalk-demo/src/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/**

钉钉的签名：https://ding-doc.dingtalk.com/doc#/faquestions/hxs5v9
签名算法为HmacSHA256，签名数据是当前时间戳timestamp，密钥是appId对应的appSecret

timestamp : 当前时间戳，单位是毫秒
appId : 扫码登录应用的appId
*/
var dingTalk = model.NewDingTalkConfig()

func Signature(timestamp, appSecret string) string {
	key := []byte(appSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(timestamp))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return url.QueryEscape(sign)
}

/**
https://ding-doc.dingtalk.com/doc#/serverapi2/kymkv6/M3fY1
服务端通过临时授权码获取授权用户的个人信息
code: 扫码获取的临时授权码
*/
func GetUserInfoByCode(code string) model.UserInfo {
	//dingTalk := model.NewDingTalkConfig()
	// todo :需要毫秒
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	//签名
	sign := Signature(timestamp, dingTalk.AppSecret)
	api := GetUserInfoByCodeUrl(dingTalk.AppId, timestamp, sign)

	data := make(map[string]string)
	data["tmp_auth_code"] = code
	args, _ := json.Marshal(data)

	resp, err := http.Post(api, "application/json", bytes.NewBuffer(args))
	if err != nil {
		log.Printf("网络请求异常%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("数据读取%s", err)
	}
	res := &model.Response{}
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Printf("数据转换异常%s", err)
	}
	if res.Code != 0 {
		log.Printf("接口请求异常%s", res.Msg)
	}
	return res.UserInfo
}

func GetUserIdByUnionId(unionId string) {
	//token := GetAccessToken()
	//_ := GetUserIdByUnionIdUrl(token, unionId)
}

/**
获取token
*/
func GetAccessToken() (string, uint16) {
	//dingTalk := model.NewDingTalkConfig()
	api := GetAccessTokenUrl(dingTalk.AgentKey, dingTalk.AgentSecret)
	resp, err := http.Get(api)
	if err != nil {
		log.Printf("网络请求异常%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("数据读取%s", err)
	}
	res := &model.Response{}
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Printf("数据转换异常%s", err)
	}
	if res.Code != 0 {
		log.Printf("接口请求异常%s", res.Msg)
	}
	return res.AccessToken, res.ExpiresIn
}
