package entities

import "time"

type ProductCore struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	UserID      uint      `json:"user_id" form:"user_id"`
	User        UserCore  `json:"user" form:"user"`
}

type ProductRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	UserID      uint   `json:"user_id" form:"user_id"`
}

type ProductResponse struct {
	Name        string       `json:"name" form:"name"`
	Description string       `json:"description" form:"description"`
	CreatedAt   time.Time    `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" form:"updated_at"`
	User        UserProResponse `json:"user" form:"user"`
}
