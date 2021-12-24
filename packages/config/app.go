package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	database *gorm.DB
)

func ConnectToDatabase() {
	db, err := gorm.Open("mysql", "root:Yuheng@0204@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	database = db
}

func GetDB() *gorm.DB {
	return database
}
