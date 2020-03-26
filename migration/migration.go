package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"setzna/db"
	"setzna/model"
)

func main() {
	db := db.Connection()
	defer db.Close()

	db.AutoMigrate(&model.Post{})
}
