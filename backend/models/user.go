// models/user.go

package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Lastname  string    `json:"lastname"`
	Firstname string    `json:"firstname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Key       string    `json:"key"`
	Provider  string    `json:"provider"`
	Dob       time.Time `json:"dob"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserProfile struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
