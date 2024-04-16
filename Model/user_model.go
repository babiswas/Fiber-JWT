package Model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string    `json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
