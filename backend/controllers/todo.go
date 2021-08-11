package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/J0Miles/todo-percipia/backend/config"
	"github.com/J0Miles/todo-percipia/backend/models"
	"github.com/gorilla/mux"
)

var (
	id         int
	title      string
	completed  bool
	created_at time.Time
	view       = template.Must(template.ParseFiles("./views/index.html"))
	database   = config.Database()
)

func Show(w http.ResponseWriter, r *http.Request) {
	// Query string to get all results
	statement, err := database.Query(`SELECT * FROM todos`)

	if err != nil {
		log.Fatal(err)
	}

	var todos []models.Todo

	// Scans for the given table columns from DB
	for statement.Next() {
		err = statement.Scan(&id, &title, &created_at, &completed)

		if err != nil {
			fmt.Println(err)
		}

		// Assigns db values to model
		todo := models.Todo{
			Id:         id,
			Title:      title,
			Updated_at: time.Now(),
			Created_at: created_at,
			Completed:  completed,
		}
		todos = append(todos, todo)
	}
	data := models.View{
		Todos: todos,
	}
	_ = view.Execute(w, data)
	fmt.Println(data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	// Attempting to insert multiple values into db exec
	title := r.FormValue("title")

	// Instantiate created_at time
	createdAt := time.Now()

	_, err := database.Exec(`INSERT INTO todos (title, createdAt) VALUES (?, ?)`, title, createdAt)

	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", 303)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// This works when using postman. I believe on redirect this is getting overwritten by cached task list? Not sure
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)

	_, err := database.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 303)
}

func Update(w http.ResponseWriter, r *http.Request) {
	// Takes in new title from form input
	newTitle := r.FormValue("newTitle")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(newTitle)

	_, err := database.Exec(`UPDATE todos SET title = ? WHERE id = ?`, newTitle, id)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", 303)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	// Sets selected task to complete
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)

	_, err := database.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 303)
}
