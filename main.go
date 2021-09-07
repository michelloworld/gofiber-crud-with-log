package main

import (
	"log"
	"time"
	"user-api/app/db"
	"user-api/app/routes"
	"user-api/app/zaplogger"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	zaplogger.Init()
	db.InitDB()

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		c.Next()
		stop := time.Now()
		latency := stop.Sub(start).Round(time.Millisecond)
		zaplogger.Log.Sugar().Infof("%v | %v | %v | %v | %v", c.Response().StatusCode(), latency, c.IP(), c.Method(), c.Path())
		return nil
	})
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
