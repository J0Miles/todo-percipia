package ui

import "github.com/J0Miles/todo-percipia/backend/models"

type Service interface {
	GetAllTodos() ([]models.Todo, error)
}
