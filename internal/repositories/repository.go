package repositories

import (
	"context"
	"time"

	"github.com/renan-parise/xcal-analytics/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *Repository) Get(hash string) (*entities.Analytics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var analytics entities.Analytics
	err := r.db.FindOne(ctx, bson.M{"hash": hash}).Decode(&analytics)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &analytics, err
}

func (r *Repository) Update(analytics *entities.Analytics) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"hash": analytics.Hash}
	update := bson.M{
		"$set": bson.M{
			"analytes":   analytics.Analytes,
			"updated_at": analytics.UpdatedAt,
		},
	}
	_, err := r.db.UpdateOne(ctx, filter, update)
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
