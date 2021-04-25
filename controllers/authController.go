package controllers

import (
	"../models"
	"../database"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err:= c.BodyParser(&data); err != nil{
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{

		Email:    data["email"],
		Password: password,
	}
	database.DB.Create(&user)

	return c.JSONP(user)
}
