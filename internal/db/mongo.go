package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type From imported Packages
type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// Creating Connection
func NewMongo(uri, dbName string) (*Mongo, error) {

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err

	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		return nil, err
	}
	log.Println("Successfully connected to MongoDB")

	return &Mongo{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

func (m *Mongo) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := m.Client.Disconnect(ctx)
	if err != nil {
		log.Println("Error Disconnecting From DB ", err)
		return err
	}
	log.Println("Successfully Disconnected From DB ")
	return nil
}
