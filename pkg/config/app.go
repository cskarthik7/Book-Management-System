package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			"root",
			"csk1999",
			"localhost",
			3306,
			"simplerest",
		),
	)
	if err != nil{
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}