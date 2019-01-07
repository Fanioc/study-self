package service

import (
	"github.com/kataras/iris"
	"log"
)

//StartWebService
func StartWebService(port string) {
	appWeb := iris.New()

	//注册路由
	RegisiterRoute(appWeb)

	//运行服务
	err := appWeb.Run(iris.Addr(":" + port))

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}

}

func StartWebScoketService(port string) {

}
