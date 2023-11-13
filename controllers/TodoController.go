package controllers

import (
	"bankcapital.co.id/tryfiber/models"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllTodo is a function to get all todo data from database
// @Summary Get all todo
// @Description Get all todo
// @Tags todos
// @Accept json
// @Produce json
// @Router /api/todos [get]
func GetAllTodos(ctx *fiber.Ctx) error {
	collection := mgm.Coll(&models.Todo{})
	todos := []models.Todo{}

	err := collection.SimpleFind(&todos, bson.D{})
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})
}

// GetTodoByID is a function to get a todo by ID
// @Summary Get todo by ID
// @Description Get todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Router /api/todos/{id} [get]
func GetTodoByID(ctx *fiber.Ctx) error {
	var id string = ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// RegisterTodo registers a new todo data
// @Summary Register a new todo
// @Description Register todo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Register todo"
// @Router /api/todos [post]
func CreateTodo(ctx *fiber.Ctx) error {
	params := new(struct {
		Title       string
		Description string
	})

	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Description) == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title or description not specified.",
		})
	}

	todo := models.CreateTodo(params.Title, params.Description)
	err := mgm.Coll(todo).Create(todo)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// UpdateTodo update todo data
// @Summary Update todo
// @Description Update todo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Update todo"
// @Router /api/todos/{id} [patch]
func ToggleTodoStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
	}

	todo.Done = !todo.Done

	err = collection.Update(todo)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// DeleteTodo function removes a todo by ID
// @Summary Remove todo by ID
// @Description Remove todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Router /api/todos/{id} [delete]
func DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
	}

	err = collection.Delete(todo)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}
