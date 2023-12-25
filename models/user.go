package models

import "gorm.io/gorm"

// struct user gorm model
type User struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string    `json:"name" form:"name"`
	Email       string    `gorm:"unique" json:"email" form:"email"`
	Password    string    `json:"password" form:"password"`
	Address     string    `json:"address" form:"address"`
	PhoneNumber string    `json:"phone_number" form:"phone_number"`
	Role        string    `json:"role" form:"role"`
	Products    []Product 
}
