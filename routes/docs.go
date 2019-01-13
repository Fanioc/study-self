package routes

import (
	"github.com/kataras/iris"
	"github.com/swaggo/swag"
)

var docs = Route{
	Name:    "Doc Root",
	Method:  "Party",
	Pattern: "/docs",
	Routes: []Route{
		{
			"Doc", // Name
			"GET", // HTTP method
			"/",   // Route pattern
			func(ctx iris.Context) {
				_ = ctx.ServeFile("./docs/swagger/index.html", false)
			},
			nil,
		},
		{
			"Doc",           // Name
			"GET",           // HTTP method
			"/swagger.json", // Route pattern
			func(ctx iris.Context) {
				doc, _ := swag.ReadDoc()
				_, _ = ctx.WriteString(doc)
			},
			nil,
		},
		{
			"Doc",       // Name
			"GET",       // HTTP method
			"/{f:path}", // Route pattern
			func(ctx iris.Context) {
				_ = ctx.ServeFile("./docs/swagger/"+ctx.Params().Get("f"), false)
			},
			nil,
		},
	},
}
