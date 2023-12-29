package repositories

import (
	"gorm-example/model"

	"gorm.io/gorm"
)

type UserRepository struct{ DB *gorm.DB }

func NewUserReopsitory(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) Create(user *model.User) *model.User {
	r.DB.Create(&user)
	return user
}
