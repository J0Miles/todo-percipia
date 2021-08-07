package usecases

import "github.com/J0Miles/todo-percipia/backend/models"

type MockTodo interface {
	GetAllTodos() ([]models.Todo, error)
}
