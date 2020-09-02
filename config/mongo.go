package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() (*mongo.Client, error) {
	host := os.Getenv("DB_HOST")

	c, e := mongo.NewClient(options.Client().ApplyURI("mongodb://" + host))
	if e != nil {
		return nil, e
	}

	if e = c.Connect(context.Background()); e != nil {
		return nil, e
	}

	return c, nil
}
