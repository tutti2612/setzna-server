package main

import (
	"setzna/internal/db"
	"setzna/internal/model"
)

func main() {
	db := db.Connection()
	defer db.Close()

	db.AutoMigrate(&model.Message{})
}
