package controllers

import (
	"fmt"
	"github.com/J0Miles/todo-percipia/backend/config"
	"github.com/J0Miles/todo-percipia/backend/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

var (
	id          int
	title       string
	description string
	completed   bool
	updated_at  time.Time
	created_at  time.Time
	view        = template.Must(template.ParseFiles("./views/index.html"))
	database    = config.Database()
)

func Show(w http.ResponseWriter, r *http.Request) {
	statement, err := database.Query(`SELECT * FROM todos`)

	if err != nil {
		fmt.Println(err)
	}

	var todos []models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &title, &created_at, &description, &completed)

		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:          id,
			Title:       title,
			Description: description,
			Updated_at:  time.Now(),
			Created_at:  time.Now(),
			Completed:   completed,
		}

		todos = append(todos, todo)
	}

	data := models.View{
		Todos: todos,
	}

	fmt.Println(data)

	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	createdAt := time.Now()

	fmt.Println(createdAt)
	fmt.Println(title)
	_, err := database.Exec(`INSERT INTO todos (title) VALUE (?)`, title)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}
