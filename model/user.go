package model

import "time"

type User struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;max=100"`
	Email     string    `gorm:"column:email;max=100"`
	Password  string    `gorm:"column:password;max=8"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}

type UserRequest struct {
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name" alidate:"required,max=100"`
	Email    string `json:"email" alidate:"required,email,max=100"`
	Password string `json:"password" alidate:"required,max=8"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updateed_at"`
}
