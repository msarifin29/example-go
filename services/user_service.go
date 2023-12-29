package services

import (
	"context"
	"fmt"
	"gorm-example/model"
	"gorm-example/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository repositories.UserRepository
	Validate       *validator.Validate
	DB             *gorm.DB
}

func NewUserService(r repositories.UserRepository, v *validator.Validate, DB *gorm.DB) *UserService {
	return &UserService{UserRepository: r, Validate: v, DB: DB}
}

func (s *UserService) Create(ctx context.Context, request model.UserRequest) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		fmt.Println("err validate :", err)
		return nil, fiber.ErrBadRequest
	}
	if err := tx.Commit().Error; err != nil {
		fmt.Println("err commit :", err)
		return nil, fiber.ErrInternalServerError
	}

	userRequest := model.User{
		ID:       uuid.NewString(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user := s.UserRepository.Create(&userRequest)
	return &model.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
