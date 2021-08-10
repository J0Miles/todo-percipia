// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/gorilla/mux"
// 	log "github.com/sirupsen/logrus"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"io"
// 	"net/http"
// )

// var db, _ = gorm.Open(mysql.Open("root:P@ssW0rd@localhost:3306/todolist?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

// type TodoItemModel struct {
// 	Id          int    `gorm:"primary_key"`
// 	Title       string `json:Title`
// 	CreatedAt   string
// 	UpdatedAt   string
// 	Description string
// 	Completed   bool
// }

// var Todos []TodoItemModel

// // var dsn = "root:P@ssW0rd@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
// // var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// func Healthz(w http.ResponseWriter, r *http.Request) {
// 	log.Info("API Health is OK")
// 	w.Header().Set("Content-Type", "application/json")
// 	io.WriteString(w, `{"alive": true}`)
// }

// func init() {
// 	log.SetFormatter(&log.TextFormatter{})
// 	log.SetReportCaller(true)
// }

// func main() {
// 	//dsn := "root:P@ssW0rd@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
// 	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	//if err != nil {
// 	//	log.Fatal(err)
// 	//}
// 	fmt.Println("DB: ", db)
// 	log.Info("Starting Todolist API server")
// 	router := mux.NewRouter()
// 	router.HandleFunc("/healthz", Healthz).Methods("GET")
// 	router.HandleFunc("/todo", CreateItem).Methods("POST")
// 	http.ListenAndServe(":8000", router)
// }

// func CreateItem(w http.ResponseWriter, r *http.Request) {
// 	description := r.FormValue("description")
// 	log.WithFields(log.Fields{"description": description}).Info("Add new TodoItem. Saving to database.")
// 	todo := &TodoItemModel{Description: description, Completed: false}
// 	db.Create(&todo)
// 	result := db.Last(&todo)
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Println("THE RESULT: ", result)
// 	json.NewEncoder(w).Encode(result)
// }

package main

import (
	"github.com/J0Miles/todo-percipia/backend/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in .env")
	}

	error := http.ListenAndServe(":"+port, routes.Init())

	if error != nil {
		log.Fatal(err)
	}
}
