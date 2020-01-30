package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	DB      *mongo.Database
	Student *StudentRepo
}

var DB *mongo.Database

func Connect(dbUrl string) (*Repository, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	defer cancel()

	conn := client.Database("interns")

	DB = conn

	repo := &Repository{
		DB:      conn,
		Student: GetStudentRepo(),
	}

	return repo, nil
}
