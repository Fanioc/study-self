package routes

import (
	"github.com/fanioc/study-self/web/controllers"
)

var studyRoutes = Route{
	Name:        "自习接口api",
	Method:      "Party",
	Pattern:     "study.", // sub domain
	HandlerFunc: nil,
	Routes: []Route{
		{
			"index",
			"GET",
			"/",
			controllers.GetStringByInt,
			nil,
		},
	},
}
