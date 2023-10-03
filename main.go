package main

import (
	"go-api/config"
	"go-api/module/user"
	"log"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	DEFAULT_PORT := config.GetEnv("PORT")
	app := fiber.New();

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())
	
	if(DEFAULT_PORT == "") {
		DEFAULT_PORT = config.DEFAULT_PORT
	}

	config.ConnectDB()
	user.UserRoutes(app)

	if err := app.Listen(":" + DEFAULT_PORT); err != nil {
		log.Fatal(err)
 	}

	app.Get("/swagger/*", swagger.Handler) // default
}