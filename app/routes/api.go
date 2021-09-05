package routes

import (
	"user-api/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteAPI(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/users", controllers.UsersController{}.Index)
	v1.Get("/users/:id", controllers.UsersController{}.Show)
	v1.Post("/users", controllers.UsersController{}.Create)
	v1.Put("/users/:id", controllers.UsersController{}.Update)
	v1.Delete("/users/:id", controllers.UsersController{}.Delete)
}
