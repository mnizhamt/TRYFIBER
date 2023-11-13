package main

import (
	"log"
	"os"

	_ "bankcapital.co.id/tryfiber/docs"
	"bankcapital.co.id/tryfiber/router"

	"github.com/Kamva/mgm/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}

	err := mgm.SetDefaultConfig(nil, "todos", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// OpenAPI documentation
	app.Get("/docs/*", swagger.HandlerDefault)

	router.SetupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}
