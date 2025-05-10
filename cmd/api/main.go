package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/internal/db"
	"example.com/internal/models"
	"example.com/internal/repository"
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

	user := models.User{
		ID:        primitive.NewObjectID(),
		Username:  "John Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = userRepo.CreateUser(context.Background(), &user)
	if err != nil {
		log.Fatal(err)
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

}
