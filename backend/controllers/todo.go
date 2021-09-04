package controllers

import (
	"encoding/json"
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
	id          int
	title       string
	description string
	completed   bool
	created_at  time.Time
	view        = template.Must(template.ParseFiles("./views/index.html"))
	database    = config.Database()
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Created_at  time.Time
	complete    bool
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	// Query string to get all results
	statement, err := database.Query(`SELECT * FROM todos`)

	if err != nil {
		log.Fatal(err)
	}

	var todos []models.Todo

	// Scans for the given table columns from DB
	for statement.Next() {
		err = statement.Scan(&id, &title, &description, &created_at, &completed)

		if err != nil {
			fmt.Println(err)
		}

		// Assigns db values to model
		todo := models.Todo{
			Id:          id,
			Title:       title,
			Description: description,
			Updated_at:  time.Now(),
			Created_at:  created_at,
			Completed:   completed,
		}
		todos = append(todos, todo)
	}
	data := models.View{
		Todos: todos,
	}
	fmt.Println(data)
	fmt.Println(todos)
	p, err := json.Marshal(todos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(p))
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
	// return data
	//	_ = view.Execute(w, data)
	//fmt.Println(data)
	// }

}

func Add(w http.ResponseWriter, r *http.Request) {
	// Attempting to insert multiple values into db exec
	// Instantiate created_at time
	timeNow := time.Now()
	fmt.Println(timeNow)
	fmt.Println("APPLES!!!!!!!!!!", r.Body)
	decoder := json.NewDecoder(r.Body)
	var t models.Todo
	err := decoder.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := database.Prepare("INSERT INTO todos(title, description, created_at) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, e := stmt.Exec(t.Title, t.Description, time.Now().Format("2006-01-02 15:04:05"))
	if e != nil {
		log.Fatal(e)
		fmt.Println(e)
	}

	response, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Println("RESPONSE!!!!!!!!!", res)

	// Allow Cross-Site Request
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	//	w.Header().Set("Content-Type", "application/json")
	//(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	//(w).Header().Set("Content-type", "application/json")
	//fmt.Println(w)
	//w.WriteHeader(http.StatusOK)
	w.Write(response)
	w.WriteHeader(200)

	fmt.Println("insert id", id)
	json.NewEncoder(w).Encode(res)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	// This works when using postman. I believe on redirect this is getting overwritten by cached task list? Not sure
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)

	_, err := database.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(200)
	http.Redirect(w, r, "/", 200)
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
