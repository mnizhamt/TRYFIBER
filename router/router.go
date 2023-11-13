package router

import (
	"bankcapital.co.id/tryfiber/controllers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("TryFiber App")
	})

	// Middleware
	api := app.Group("/api", logger.New())
	auth := api.Group("/auth")
	todo := api.Group("/todos")

	// Auth
	auth.Post("/login", controllers.UserSignIn)
	todo.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// Todos
	todo.Get("/", controllers.GetAllTodos)
	todo.Get("/:id", controllers.GetTodoByID)
	todo.Post("", controllers.CreateTodo)
	todo.Patch("/:id", controllers.ToggleTodoStatus)
	todo.Delete("/:id", controllers.DeleteTodo)

	// User
	// user := api.Group("/user")
	// user.Get("/:id", handler.GetUser)
	// user.Post("/", handler.CreateUser)
	// user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	// product := api.Group("/product")
	// product.Get("/", handler.GetAllProducts)
	// product.Get("/:id", handler.GetProduct)
	// product.Post("/", middleware.Protected(), handler.CreateProduct)
	// product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
