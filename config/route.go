package config

import (
	"gorm-example/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteCofig struct {
	App            *fiber.App
	UserController *controller.UserController
}

func (r *RouteCofig) Setup() {
	r.App.Post("/api/users", r.UserController.Create)
	r.App.Get("/api/users", r.UserController.FindAll)
}
