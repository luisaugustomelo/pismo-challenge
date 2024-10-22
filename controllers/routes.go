package controllers

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	SetupSwaggerRoutes(app)
}
