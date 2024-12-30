package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Client, error) {
	url := os.Getenv("URL")

	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return client, nil
}
