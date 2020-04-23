package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
)

func main() {
	cfg, err := ini.Load("./config/cfg.ini")

	if err != nil {
		log.Fatal("load dingtalk error", err)
	}

	rout := gin.Default()
	rout.LoadHTMLGlob("./view/**")
	//rout.Static("static","./static")
	rout.StaticFile("/favicon.ico", "./favicon.ico")

	rout.GET("/", func(context *gin.Context) {
		redirect := cfg.Section("dingtalk").Key("goto").String()
		context.HTML(http.StatusOK, "index.html", gin.H{"title": "钉钉扫码登录", "redirect": redirect})
	})

	/**
	钉钉扫码登录
	*/
	rout.GET("/scan/login", func(cxt *gin.Context) {
		code := cxt.Query("code")
		log.Println(code)

		cxt.HTML(http.StatusOK, "succeed.html", gin.H{"title": "登录成功"})
	})
	rout.Run(":8100")
}
