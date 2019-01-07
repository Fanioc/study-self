package routes

import (
	"github.com/kataras/iris"
	"strconv"
)

var test = Route{
	Name:        "测试分组",  // Name
	Method:      "Party", // HTTP method
	Pattern:     "/test", // Route pattern
	HandlerFunc: nil,
	Routes: Routes{
		{
			"GetAccount",   // Name
			"GET",          // HTTP method
			"/{accountId}", // Route pattern
			func(ctx iris.Context) { //对应一个方法
				accountId, err := ctx.Params().GetInt("accountId")
				if err != nil {
					_, _ = ctx.Writef("error while trying to parse userid parameter," +
						"this will never happen if :int is being used because if it's not integer it will fire Not Found automatically.")
					ctx.StatusCode(iris.StatusBadRequest)
					return
				}
				_, _ = ctx.Write([]byte(`{"result":"OK"}`))
				accountIds := strconv.FormatInt(int64(accountId), 10)
				_, _ = ctx.WriteString(accountIds)
			},
			nil,
		},
		{
			"GetAccount", // Name
			"GET",        // HTTP method
			"/",          // Route pattern
			func(ctx iris.Context) { //对应一个方法
				_, _ = ctx.Write([]byte(`{"result":"index"}`))
			},
			nil,
		},
	},
}
