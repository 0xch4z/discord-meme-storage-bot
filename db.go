package main

import (
	"fmt"
	l "log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	uri := fmt.Sprintf("root:%s@db/dmsb?parseTime=true&charset=utf8", conf.DB.RootPassword)
	database, err := gorm.Open("mysql", uri)
	if err != nil {
		l.Fatal("could not open database connection", err.Error())
	}
	database.AutoMigrate(&meme{}, &guild{})
	db = database
}
