package repository

import (
	"context"
	"time"

	"example.com/internal/db"
	"example.com/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
	CreateTodo(todo *models.Todo) (*models.Todo, error)
	GetTodo(id string) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) (*models.Todo, error)
	DeleteTodo(id string) error
}

type MongoTodoRepository struct {
	db         *db.Mongo
	collection *mongo.Collection
}

func NewMongoTodoRepository(db *db.Mongo) *MongoTodoRepository {
	return &MongoTodoRepository{
		db:         db,
		collection: db.Database.Collection("todos"),
	}
}

func (r *MongoTodoRepository) CreateTodo(ctx context.Context, todo *models.Todo) error {

	todo.ID = primitive.NewObjectID()

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, todo)
	return err
}

func (r *MongoTodoRepository) DeleteTodo(ctx context.Context, userid string) error {
	filter := bson.M{"userid": userid}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err

}

func (r *MongoTodoRepository) GetTodo(ctx context.Context, userid string) (*models.Todo, error) {
	filter := bson.M{"userid": userid}
	var todo models.Todo
	err := r.collection.FindOne(ctx, filter).Decode(&todo)
	return &todo, err
}
