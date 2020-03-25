package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func migration() {
	db, err := gorm.Open("mysql", "setzna:setzna@(mysql)/setzna?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	defer db.Close()

	db.AutoMigrate(&Post{})
}
