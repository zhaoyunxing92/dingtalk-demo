package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing92/dingtalk-demo/config"
	"github.com/zhaoyunxing92/dingtalk-demo/sdk/dingtalk"
	"github.com/zhaoyunxing92/dingtalk-demo/services"
	"net/http"
)

type UserController struct {
	// 用户服务
	userService services.IUserService
}

func NewUserController(service services.IUserService) *UserController {
	return &UserController{service}
}

/**
钉钉扫码登录
*/
func (ctr *UserController) ScanLogin(ctx *gin.Context) {
	talkConfig := config.NewDingTalkConfig()
	//钉钉获取用户信息
	client := dingtalk.NewDingTalkScanClient(talkConfig.AppId, talkConfig.AppSecret)
	//临时授权码
	code := ctx.Query("code")
	userInfo := client.GetUserInfoByCode(code)
	//根据用户unionid获取用户
	user, err := ctr.userService.SelectByUnionId(userInfo.Unionid)

	if err != nil {
		if err.Error() == "not fount" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2,
				"msg":  nil,
				"data": userInfo,
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 3,
			"msg":  err.Error(),
			"data": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  nil,
		"data": user.Id,
	})
}

/**
根据id获取用户信息
*/
func (ctr *UserController) SelectById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := ctr.userService.SelectById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "获取用户成功",
			"data": user,
		})
	}
}
