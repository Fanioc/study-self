package main

import (
	"github.com/fanioc/study-self/service"
	_ "github.com/kataras/iris"
)

func main() {
	service.StartWebService("8080")
}
