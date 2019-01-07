package routes

import "github.com/kataras/iris"

var studyRoutes = Route{
	Name:        "自习接口api",
	Method:      "Party",
	Pattern:     "study.", // sub domain
	HandlerFunc: nil,
	Routes: Routes{
		{
			"index",
			"GET",
			"/",
			func(ctx iris.Context) { //对应一个方法
				_, _ = ctx.Write([]byte(`{"result":"study"}`))
			},
			nil,
		},
	},
}
