package model

import (
	"setzna/internal/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Message struct {
	gorm.Model
	Type      string `json:"type"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (m *Message) Save() {
	db := db.Connection()
	defer db.Close()

	db.Create(m)
}
