package usecases_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/J0Miles/todo-percipia/backend/models"
	"github.com/J0Miles/todo-percipia/backend/usecases"
)

var dummyTodos = []models.Todo{
	{
		Title: "Testing",
		Description: "Testing",
		Created_at: "Now",
		Updated_at: "Last present moment",
		Completed: true,
	},
	{
		Title: "Testing",
		Description: "Testing",
		Created_at: "Now",
		Updated_at: "Last present moment",
		Completed: false,
	},
	{
		Title: "Testing",
		Description: "Testing",
		Created_at: "",
		Updated_at: "",
		Completed: true,
	},
}

type MockTodos struct{}
type InvalidMockTodo struct{}

func (InvalidMockTodo) GetAllTodos() ([]models.Todo, error) {
	return nil, fmt.Errorf("Something went wrong")
}

func (MockTodos) GetAllTodos() ([]models.Todo, error) {
	return dummyTodos, nil
}

func TestGetTodos(t *testing.T) {

	// Test
	t.Run("Returns ErrInteral when MockTodo returns error", func(t *testing.T) {
		expect := expectate.Expect(t)

		mock := new(InvalidMockTodo)

		todos, err := usecases.GetTodos(mock)

		expect(err).ToBe(usecases.ErrInteral)

		if todos != nil {
			t.Fatalf("Expected todos to be nil; Got: %v", todos)
		}
	})

	// Test
	t.Run("Returns todos from MockTodo", func(t *testing.T) {
		expect := expectate.Expect(t)

		mock := new(MockTodos)

		todos, err := usecases.GetTodos(mock)

		expect(err).ToBe(nil)
		expect(todos).ToEqual(dummyTodos)
	})
}
