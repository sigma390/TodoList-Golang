package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/internal/api"
	"example.com/internal/db"

	"example.com/internal/repository"
	"github.com/go-chi/chi/v5"
)

func main() {

	//setup Db
	mongo, err := db.NewMongo("mongodb://localhost:27017", "todo")
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Close()

	//create a Repo
	todoRepo := repository.NewMongoTodoRepository(mongo)
	//get Handlers
	todoHandler := api.NewTodoHandler(todoRepo)

	//Create A Router
	mux := chi.NewRouter()
	home := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Todo API"))
	}

	mux.Post("/todos", todoHandler.CreateTodo)
	mux.Get("/todos/{id}", todoHandler.GetTodo)
	mux.Delete("/todos/{id}", todoHandler.DeleteTodo)
	mux.Get("/", home)

	//start server
	fmt.Println("Server starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", mux))

}
