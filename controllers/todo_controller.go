package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	model "sample-golang-fiber-started/models"
	"strconv"
)

var todos = []*model.Todo{
	{Id: 1, Title: "Learn Golang", Completed: false},
	{Id: 2, Title: "Build RESTful API in Go", Completed: false},
}

func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}

func CreateTodo(c *fiber.Ctx) error {
	var body model.Request
	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	todo := &model.Todo{
		Id:        len(todos) + 1,
		Title:     body.Title,
		Completed: false,
	}

	todos = append(todos, todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func GetTodo(c *fiber.Ctx) error {

	paramId := c.Params("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
		})
	}

	for _, todo := range todos {
		if todo.Id == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"todo": todo,
				},
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	// find parameter
	paramId := c.Params("id")

	// convert parameter string to int
	id, err := strconv.Atoi(paramId)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}

	var body model.RequestPoint
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	var todo *model.Todo

	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}

	if todo.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
		})
	}

	if body.Title != nil {
		todo.Title = *body.Title
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	paramId := c.Params("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}

	for i, todo := range todos {
		if todo.Id == id {

			todos = append(todos[:i], todos[i+1:]...)

			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
				"success": true,
				"message": "Deleted Succesfully",
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}
