package test

import (
	"gorm-example/config"
	"gorm-example/controller"
	"gorm-example/repositories"
	"gorm-example/services"

	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	fiberConfig := config.NewFiber()
	v := config.NewValidator()
	db := config.Connection()
	r := repositories.NewUserReopsitory(db)
	service := services.NewUserService(*r, v, db)
	controller := controller.NewUserController(service)
	route := config.RouteCofig{App: fiberConfig, UserController: controller}
	route.Setup()
	return fiberConfig
}
