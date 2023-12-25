package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	UserID      uint   `gorm:"foreignKey:UserID" json:"user_id" form:"user_id"`
	User        User   
}
