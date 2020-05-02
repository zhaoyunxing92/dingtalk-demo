package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing92/dingtalk-demo/config"
	"net/http"
	"net/url"
)

type DingtalkController struct {
}

func NewDingtalkController() *DingtalkController {
	return &DingtalkController{}
}

/**
获取钉钉扫码配置
*/
func (ctr *DingtalkController) GetDingtalkScanCfg(ctx *gin.Context) {
	talkConfig := config.NewDingTalkConfig()
	// todo 扫码的url顺序不能乱
	u := fmt.Sprintf("%s/?appid=%s&response_type=code&scope=snsapi_login&state=STATE&redirect_uri=%s", talkConfig.ScanAuthorizeUrl, talkConfig.AppId, talkConfig.ScanRedirectUrl)
	data := make(map[string]string)
	data["url"] = u
	data["goto"] = url.QueryEscape(u)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取钉钉扫码配置成功",
		"data": data,
	})
}
