package main

import (
	"log"
	"user-api/app/db"
	"user-api/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(fiber.Map{"error": fiber.Map{"message": r}})
			}
		}()
		return c.Next()
	})

	routes.RouteAPI(app)

	log.Fatal(app.Listen(":3001"))
}
