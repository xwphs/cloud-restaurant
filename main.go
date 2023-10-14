package main

import (
	"log"
	"xwphs/cloud-restaurant/controller"
	"xwphs/cloud-restaurant/tool"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := tool.ParseConfig("config/config.json")
	if err != nil {
		log.Fatal("Parse config file error: ", err)
	}
	engine := gin.Default()
	RegisterRouter(engine)
	engine.Run(config.AppHost + ":" + config.AppPort)
}

// 注册路由
func RegisterRouter(engine *gin.Engine) {
	new(controller.HelloController).Router(engine)
}
