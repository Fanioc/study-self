package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var Err error

func ConnectDb(config DbConfig) {

	Db, Err = gorm.Open("mysql", config.User+":"+config.Password+"@tcp(127.0.0.1:3306)/"+config.Dbname+"?charset="+config.Charset+"&parseTime=True&loc=Local")

	if Err != nil {
		panic("connect databases faile:" + Err.Error())
	}

	Db.DB().SetMaxOpenConns(2000)
	Db.DB().SetMaxIdleConns(1000)
}
