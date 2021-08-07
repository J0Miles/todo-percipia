package usecases

import "github.com/J0Miles/todo-percipia/backend/models"

func GetTodos(mock MockTodo) ([]models.Todo, error) {
	todos, err := mock.GetAllTodos()
	if err != nil {
		return nil, ErrInteral
	}

	return todos, nil
}
