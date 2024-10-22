package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/luisaugustomelo/pismo-challenge/db"
	"github.com/luisaugustomelo/pismo-challenge/migrations"
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/luisaugustomelo/pismo-challenge/controllers"
	"github.com/luisaugustomelo/pismo-challenge/utils/config"
)

func main() {
	db.Connect()
	defer db.CloseConnection()

	// Run migrations
	migrations.Migrate()

	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	// disable some sensitive headers to avoid some attacks, e.g: csrf and XSS
	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	config.LoadEnv()

	// Middleware
	app.Use(logger.New())
	controllers.SetupRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		log.Fatalf("couldn't listen to port %s \n error: %s", config.PORT, err.Error())
	}
}
