package controllers

import (
	"github.com/fanioc/study-self/config"
	"github.com/fanioc/study-self/database"
	"github.com/kataras/iris"
	"log"
)

// @Summary 测试接口ffff
// @Description 描述
// @ID get-string-by-int
// @Accept  json
// @tags V1
// @Produce  json
// @Param   some_id      path   int     true  "Some ID" Format(int64)
// @Param   some_id      body controllers.Pet3 true  "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} controllers.Pet3 "We need ID!!"
// @Failure 404 {object} routes.Route "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(c iris.Context) {
	results := make([]Result, 0)
	database.Db.Raw("SELECT name, age FROM users ").Find(&results)

	config.Conf = config.New()
	driver := config.Conf.Get("database.local")
	log.Println(driver)
	_, _ = c.Writef("%d", results[1].Age)
	log.Println(results)
}

type Result struct {
	Name string
	Age  int
}

// @Description 测试接口
// @ID get-struct-array-by-string
// @tags V2
// @Accept  json
// @Produce  json
// @Param some_id path string true "Some ID"
// @Param category query int true "Category" Enums(1, 2, 3)
// @Param offset query int true "Offset" Mininum(0) default(0)
// @Param limit query int true "Limit" Maxinum(50) default(10)
// @Param q query string true "q" Minlength(1) Maxlength(50) default("")
// @Success 200 {string} string	"ok"
// @Failure 400 {object} routes.Route "We need ID!!"
// @Failure 404 {object} routes.Route "Can not find ID"
// @Security ApiKeyAuth
// @Security BasicAuth
// @Security OAuth2Application[write]
// @Security OAuth2Implicit[read, admin]
// @Security OAuth2AccessCode[read]
// @Security OAuth2Password[admin]
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(c iris.Context) {
	//write your code
}

// @Summary Upload file
// @Description Upload file
// @ID file.upload
// @tags V3
// @Accept  multipart/form-data
// @Produce  json
// @Param   file formData file true  "this is a test file"
// @Success 200 {string} string "ok"
// @Failure 400 {object} routes.Route "We need ID!!"
// @Failure 404 {object} routes.Route "Can not find ID"
// @Router /file/upload [post]
func Upload(ctx iris.Context) {
	//write your code
}

// @Summary use Anonymous field
// @Success 200 {object} routes.Route "ok"
func AnonymousField() {

}

// @Summary use pet2
// @Success 200 {object} routes.Route "ok"
func Pet2() {

}

type Pet3 struct {
	ID int `json:"id"`
}
