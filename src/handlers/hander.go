package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing90/dingtalk-demo/src/model"
	"github.com/zhaoyunxing90/dingtalk-demo/src/tools/dingtalk"
	"log"
	"net/http"
	"net/url"
)

func Home(cxt *gin.Context) {

	dingTalk := model.NewDingTalkConfig()
	// todo 扫码的url顺序不能乱
	u := fmt.Sprintf("%s/?appid=%s&response_type=code&scope=snsapi_login&state=STATE&redirect_uri=%s", dingTalk.ScanAuthorizeUrl, dingTalk.AppId, dingTalk.ScanRedirectUrl)

	cxt.HTML(http.StatusOK, "index.html", gin.H{"title": "钉钉扫码登录",
		"goto": url.QueryEscape(u),
		"url":  u})
}

/**
code码登录
*/
func CodeLogin(cxt *gin.Context) {
	//return func(cxt *gin.Context) {
	code := cxt.Query("code")
	log.Println(code)
	dingtalk.GetUserInfoByCode(code)
	cxt.HTML(http.StatusOK, "succeed.html", gin.H{"title": "登录成功"})
	//}
}
