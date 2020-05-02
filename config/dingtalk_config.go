package config

import (
	"gopkg.in/ini.v1"
	"log"
)

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
	cfg, err := ini.Load("d:/github/go/dingtalk-demo/cfg.ini")
	if err != nil {
		log.Fatal("load cfg.ini error", err)
	}
	config := &dingTalkConfig{}
	err = cfg.Section("dingtalk").MapTo(config)
	if err != nil {
		log.Fatal("conversion DingTalkConfig error", err)
	}
	return config
}