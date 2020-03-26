package model

import (
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
