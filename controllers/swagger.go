package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	_ "github.com/luisaugustomelo/pismo-challenge/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupSwaggerRoutes(app *fiber.App) {
	app.Static("/swagger", "./swagger-ui")
	app.Get("/swagger/*", adaptor.HTTPHandler(httpSwagger.Handler()))
}
