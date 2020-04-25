package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing90/dingtalk-demo/src/handlers"
	"github.com/zhaoyunxing90/dingtalk-demo/src/model"
	"net/http"
)

func main() {

	rout := gin.Default()
	rout.LoadHTMLGlob("./view/**")
	//rout.Static("static","./static")
	rout.StaticFile("/favicon.ico", "./favicon.ico")

	/**
	  首页
	*/
	rout.GET("/config", func(context *gin.Context) {
		config := model.NewDingTalkConfig()
		context.JSON(http.StatusOK, gin.H{"config": config})
	})

	/**
	  首页
	*/
	rout.GET("/", handlers.Home)
	/**
	钉钉扫码登录
	*/
	rout.GET("/scan/login", handlers.CodeLogin)

	_ = rout.Run(":8100")
}
