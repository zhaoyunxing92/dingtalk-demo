package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zhaoyunxing92/dingtalk-demo/sdk/dingtalk/encrypt"
	"github.com/zhaoyunxing92/dingtalk-demo/sdk/dingtalk/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type IDingTalkScanClient interface {
	/**
	https://ding-doc.dingtalk.com/doc#/serverapi2/kymkv6/M3fY1
	服务端通过临时授权码获取授权用户的个人信息
	code: 扫码获取的临时授权码
	*/
	GetUserInfoByCode(code string) model.UserInfo
}

type DingTalkScanClient struct {
	/**
	扫码使用的appid
	*/
	appId,
	/**
	扫码使用秘钥
	*/
	appSecret string
}

/**
扫码客户端
*/
func NewDingTalkScanClient(appId, appSecret string) IDingTalkScanClient {
	return &DingTalkScanClient{appId, appSecret}
}

/**
根据code码获取用户
 */
func (client *DingTalkScanClient) GetUserInfoByCode(code string) model.UserInfo {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	sign := encrypt.Signature(timestamp, client.appSecret)
	//请求url
	uri := fmt.Sprintf("https://oapi.dingtalk.com/sns/getuserinfo_bycode?accessKey=%s&timestamp=%s&signature=%s", client.appId, timestamp, sign)

	// 设置参数
	data := make(map[string]string)
	data["tmp_auth_code"] = code
	args, _ := json.Marshal(data)

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(args))
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
