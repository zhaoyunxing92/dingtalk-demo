package dingtalk

import "fmt"

var host = "https://oapi.dingtalk.com"

/**
https://ding-doc.dingtalk.com/doc#/serverapi2/kymkv6/M3fY1
服务端通过临时授权码获取授权用户的个人信息
*/
func GetUserInfoByCodeUrl(appId, timestamp, sign string) string {

	return fmt.Sprintf(host+"/sns/getuserinfo_bycode?accessKey=%s&timestamp=%s&signature=%s", appId, timestamp, sign)
}

/**
https://ding-doc.dingtalk.com/doc#/serverapi2/ege851/602f4b15
根据unionid获取userid
https://oapi.dingtalk.com/user/getUseridByUnionid?access_token=ACCESS_TOKEN&unionid=xxx
*/
func GetUserIdByUnionIdUrl(token, unionId string) string {
	return fmt.Sprintf(host+"/user/getUseridByUnionid?access_token=%s&unionid=%s", token, unionId)
}

/**
获取access_token
https://oapi.dingtalk.com/gettoken?appkey=key&appsecret=secret
*/
func GetAccessTokenUrl(agentKey, agentSecret string) string {

	return fmt.Sprintf(host+"/gettoken?appkey=%s&appsecret=%s", agentKey, agentSecret)
}
