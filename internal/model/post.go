package model

import (
	"setzna/internal/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	gorm.Model
	PostType  string `json:"type"`
	Name      string `json:"name"`
	Message   string `json:"message"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (p *Post) Save() {
	db := db.Connection()
	defer db.Close()

	db.Create(p)
}
