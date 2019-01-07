package service

import (
	"github.com/fanioc/study-self/routes"
	"github.com/kataras/iris"
)

func RegisiterRoute(app *iris.Application) {

	// Iterate over the routes we declared in routes/routes_main.go and attach them to the router instance
	for _, route := range routes.Routers {
		regisiterRouter(app, route)
	}

}

// regisiter routees in Application
func regisiterRouter(app iris.Party, route routes.Route) {
	if route.Method == "Party" {

		// judge exiest middleware,protecet from err
		appParty := *new(iris.Party)
		if route.HandlerFunc == nil {
			appParty = app.Party(route.Pattern)
		} else {
			appParty = app.Party(route.Pattern, route.HandlerFunc)
		}

		for _, routeItem := range route.Routes {
			regisiterRouter(appParty, routeItem)
		}
	} else {
		app.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}

}
