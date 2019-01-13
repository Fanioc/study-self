package routes

import (
	"github.com/kataras/iris"
)

/**
 * 添加路由文件
 */
var routes = []Route{
	// study MiniApp api
	studyRoutes,

	// test Routes
	test,

	// api docs
	docs,
}

/* Router function */
func Regisiter(app *iris.Application) {

	for _, route := range routes {
		regisiterRoutes(app, route)
	}

}

// regisiter routees in Application
func regisiterRoutes(app iris.Party, route Route) {
	if route.Method == "Party" {
		// judge exiest middleware,protecet from err
		appParty := *new(iris.Party)
		if route.HandlerFunc == nil {
			appParty = app.Party(route.Pattern)
		} else {
			appParty = app.Party(route.Pattern, route.HandlerFunc)
		}
		for _, routeItem := range route.Routes {
			regisiterRoutes(appParty, routeItem)
		}
	} else {
		app.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

// Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string //
	Method      string // if Method=Party , Routes!=nil
	Pattern     string
	HandlerFunc iris.Handler
	Routes      []Route
}

// Defines the type Routes which is just an array (slice) of Route structs.
//type Routes []Route
