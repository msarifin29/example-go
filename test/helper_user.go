package test

import (
	"gorm-example/config"
	"gorm-example/model"

	"github.com/google/uuid"
)

func CreateUser() {
	db := config.Connection()

	user := model.User{
		ID:       uuid.NewString(),
		Name:     "Layla",
		Email:    "layla@mail.com",
		Password: "12345678",
	}

	db.Create(&user)
}
