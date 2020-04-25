package model

import (
	"gopkg.in/ini.v1"
	"log"
)

/**
配置文件
*/
type dingTalkConfig struct {
	/**
	扫码使用的appid
	*/
	AppId,
	/**
	扫码使用秘钥
	*/
	AppSecret,
	/**
	扫码回调地址 跟后台配置的要应用
	*/
	ScanRedirectUrl,
	/**
	扫码认证地址
	*/
	ScanAuthorizeUrl,
	/**
	微应用key
	*/
	AgentKey,
	/**
	微应用秘钥
	*/
	AgentSecret string
}

/**
获取配置信息
*/
func NewDingTalkConfig() *dingTalkConfig {
	// todo: 路径有问题
	cfg, err := ini.Load("d:/github/go/dingtalk-demo/config/cfg.ini")
	if err != nil {
		log.Fatal("load cfg.ini error", err)
	}
	config := &dingTalkConfig{}
	err = cfg.Section("DingTalk").MapTo(config)
	if err != nil {
		log.Fatal("conversion DingTalkConfig error", err)
	}
	return config
}

type UserInfo struct {
	/**
	用户在钉钉上面的昵称
	*/
	Nick string `json:"nick"`
	/**
	用户在当前开放应用内的唯一标识
	*/
	Openid string `json:"openid"`
	/**
	用户在当前开放应用所属企业内的唯一标识
	*/
	Unionid string `json:"unionid"`
}

type AccessToken struct {
}

/**
钉钉返回结果  "access_token": "fw8ef8we8f76e6f7s8df8s"
*/
type Response struct {
	Code        int      `json:"errcode"`
	Msg         string   `json:"errmsg"`
	UserInfo    UserInfo `json:"user_info"`
	AccessToken string   `json:"access_token"`
	ExpiresIn   uint16     `json:"expires_in"`
}
