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
		log.Fatal(error)
	}
}
