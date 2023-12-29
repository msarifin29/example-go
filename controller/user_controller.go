package controller

import (
	"fmt"
	"gorm-example/model"
	"gorm-example/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(u *services.UserService) *UserController {
	return &UserController{UserService: u}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	request := new(model.UserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		fmt.Println("error parser :", err)
		return fiber.ErrBadRequest
	}
	response, err := c.UserService.Create(ctx.Context(), *request)
	if err != nil {
		fmt.Println("error create controller :", err)
		return fiber.ErrInternalServerError
	}
	return ctx.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "Success create user",
		Data:   response})
}

func (c *UserController) FindAll(ctx *fiber.Ctx) error {
	users, err := c.UserService.FindAll(ctx.Context())
	if err != nil {
		fmt.Println("error find all controller :", err)
		return fiber.ErrInternalServerError
	}
	return ctx.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "Success get users",
		Data:   users,
	})
}
