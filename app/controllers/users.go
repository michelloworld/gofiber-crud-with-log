package controllers

import (
	"user-api/app/db"
	"user-api/app/models"
	"user-api/app/zaplogger"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct{}

func (UsersController) Index(c *fiber.Ctx) error {
	zaplogger.Log.Info("=== Finding all users")
	var users []models.User
	if result := db.DB.Find(&users); result.Error != nil {
		panic(result.Error.Error())
	}
	return c.JSON(users)
}

func (UsersController) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		panic(result.Error.Error())
	}
	return c.JSON(user)
}

func (UsersController) Create(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		panic(err.Error())
	}
	if result := db.DB.Create(&user); result.Error != nil {
		panic(result.Error.Error())
	}
	c.SendStatus(201)
	return c.JSON(user)
}

func (UsersController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	updateUser := models.User{}
	if err := c.BodyParser(&updateUser); err != nil {
		panic(err.Error())
	}
	findUser := models.User{}
	if result := db.DB.First(&findUser, id); result.Error != nil {
		panic(result.Error.Error())
	}
	if result := db.DB.Model(&findUser).Updates(updateUser); result.Error != nil {
		panic(result.Error.Error())
	}
	return c.JSON(updateUser)
}

func (UsersController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	findUser := models.User{}
	if result := db.DB.First(&findUser, id); result.Error != nil {
		panic(result.Error.Error())
	}
	if result := db.DB.Delete(findUser); result.Error != nil {
		panic(result.Error.Error())
	}
	return c.JSON(findUser)
}
