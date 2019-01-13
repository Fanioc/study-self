package service

import (
	"github.com/fanioc/study-self/config"
	"github.com/fanioc/study-self/database"
	"github.com/fanioc/study-self/routes"
	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"
	"log"
)

//StartWebService
func StartWebService(port string) {

	app := iris.New()

	//注册路由
	routes.Regisiter(app)

	//注册数据库实例
	DBconfig := database.DbConfig{}
	_ = config.Get("database.local").(*toml.Tree).Unmarshal(&DBconfig)
	database.ConnectDb(DBconfig)

	//注册静态文件页面
	app.StaticWeb("/", "./web/public")

	//运行服务
	err := app.Run(iris.Addr(":" + port))

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}

}

func StartWebScoketService(port string) {

}
