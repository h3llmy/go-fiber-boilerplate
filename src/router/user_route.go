package router

import (
	"app/src/controllers"
	m "app/src/middleware"
	"app/src/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(v1 fiber.Router, u services.UserService, t services.TokenService) {
	userController := controllers.NewUserController(u, t)

	user := v1.Group("/users")

	user.Get("/", m.Auth(u, "getUsers"), userController.GetUsers)
	user.Post("/", m.Auth(u, "manageUsers"), userController.CreateUser)
	user.Get("/:userId", m.Auth(u, "getUsers"), userController.GetUserByID)
	user.Patch("/:userId", m.Auth(u, "manageUsers"), userController.UpdateUser)
	user.Delete("/:userId", m.Auth(u, "manageUsers"), userController.DeleteUser)
}
