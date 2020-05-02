package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing92/dingtalk-demo/backend/web/controller"
	"github.com/zhaoyunxing92/dingtalk-demo/config"
	"github.com/zhaoyunxing92/dingtalk-demo/repositories"
	"github.com/zhaoyunxing92/dingtalk-demo/services"
	"log"
)

// web入口
func main() {
	rout := gin.Default()
	rout.LoadHTMLGlob("./view/**")
	rout.StaticFile("/favicon.ico", "./favicon.ico")

	mongo := config.NewMongoManger("mongodb://localhost:27017", "cherry")
	coll, err := mongo.GetCollection("user")
	if err != nil {
		log.Fatal(err)
	}
	//用户控制器
	userRepository := repositories.NewUserRepository(coll)
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	// 钉钉控制器
	dingtalkController := controller.NewDingtalkController()

	/**
	钉钉扫码配置
	*/
	rout.GET("/dd/scan/cfg", dingtalkController.GetDingtalkScanCfg)
	/**
	钉钉扫码登录
	*/
	rout.GET("/scan/login", userController.ScanLogin)

	/**
	根据unionId获取用户信息
	*/
	rout.GET("/user/:id", userController.SelectById)

	_ = rout.Run(":8100")
}
