package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/internal/api"
	"example.com/internal/db"
	"example.com/internal/models"
	"example.com/internal/repository"
	"github.com/go-chi/chi/v5"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	userRepo := repository.NewMongoUserRepository(mongo)

	//get Handlers
	todoHandler := api.NewTodoHandler(todoRepo)

	user := models.User{
		ID:        primitive.NewObjectID(),
		Username:  "omakr patil",
		Email:     "omakrpatil@gmail.com",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users := []models.User{{
		ID:        primitive.NewObjectID(),
		Username:  "omakr patil",
		Email:     "omakrpatil@gmail.com",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, {
		ID:        primitive.NewObjectID(),
		Username:  "anish patil",
		Email:     "anishpatil@gmail.com",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, {
		ID:        primitive.NewObjectID(),
		Username:  "Jessie patil",
		Email:     "jessiepatil@gmail.com",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}}

	err = userRepo.CreateUser(context.Background(), &user)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		err := userRepo.CreateUser(context.Background(), &user)
		if err != nil {
			log.Fatal(err)
		}
	}

	allUsers, err := userRepo.GetAllUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range allUsers {
		fmt.Println(user)
		fmt.Println("--------------------------------")
	}

	todo := models.Todo{
		ID:          primitive.NewObjectID(),
		Title:       "Buy groceries",
		Description: "Buy groceries",
		UserID:      "1",
		Status:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = todoRepo.CreateTodo(context.Background(), &todo)
	if err != nil {
		log.Fatal(err)
	}

	// err = todoRepo.DeleteTodo(context.Background(), "1")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	todo1, err := todoRepo.GetTodo(context.Background(), "1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(todo1)
	//Create A Router
	mux := chi.NewRouter()

	mux.Post("/todos", todoHandler.CreateTodo)
	mux.Delete("/todos/{id}", todoHandler.DeleteTodo)

	//start server
	http.ListenAndServe(":8080", mux)
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
