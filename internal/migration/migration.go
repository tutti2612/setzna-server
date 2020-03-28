package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"setzna/internal/db"
	"setzna/internal/model"
)

func main() {
	db := db.Connection()
	defer db.Close()

	db.AutoMigrate(&model.Message{})
}
