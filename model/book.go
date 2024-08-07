package model

import (
	"gorm.io/gorm"
)

type Book struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}
