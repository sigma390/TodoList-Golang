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

func (r *MongoUserRepository) DeleteUser(ctx context.Context, userid string) error {
	filter := bson.M{"ID": userid}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err

}

func (r *MongoUserRepository) GetUser(ctx context.Context, userid string) (*models.User, error) {
	filter := bson.M{"ID": userid}
	var user models.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	return &user, err
}

// gert All users
func (r *MongoUserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	//create Empty Filter
	filter := bson.M{}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	//defer close cursor
	defer cursor.Close(ctx)

	var users []models.User
	//decode all users
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil

}
