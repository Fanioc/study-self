package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDb(config DbConfig) (Db *gorm.DB, err error) {
	Db, err = gorm.Open("mysql", config.Dbname+":"+config.Password+"@/"+config.Dbname+"?charset="+config.Charset+"&parseTime=True&loc=Local")
	if err != nil {
		return
	}

	return
}
