package repository

import (
	"context"
	"time"

	"example.com/internal/db"
	"example.com/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	db         *db.Mongo
	collection *mongo.Collection
}

func NewMongoUserRepository(db *db.Mongo) *MongoUserRepository {
	return &MongoUserRepository{
		db:         db,
		collection: db.Database.Collection("users"),
	}
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}

func (r *MongoUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}
