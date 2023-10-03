package main

import (
	"go-api/config"
	"go-api/module/user"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

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

	app.Get("/swagger/*", swagger.HandlerDefault) // default
}