package repositories

import (
	"context"
	"time"

	"github.com/renan-parise/xcal-analytics/internal/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Collection
}

func NewAnalyticsRepository(db *mongo.Database) *Repository {
	return &Repository{db: db.Collection("analytics")}
}

func (r *Repository) Insert(analytics *entities.Analytics) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.db.InsertOne(ctx, analytics)
	return err
}

func Connect() *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	db := client.Database("xcal")
	return db
}
