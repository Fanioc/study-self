// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/

package main

import (
	_ "github.com/fanioc/study-self/docs"
	"github.com/fanioc/study-self/service"
	_ "github.com/kataras/iris"
)

//go:generate swag init
//go:generate go build
//go:generate ./study-self
func main() {
	service.StartWebService("8080")
}
