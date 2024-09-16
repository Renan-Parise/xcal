package repositories

import (
	"context"
	"time"

	"github.com/renan-parise/gofreela/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Collection
}

func NewJobsRepository(db *mongo.Database) *Repository {
	return &Repository{db: db.Collection("jobs")}
}

func (r *Repository) Insert(jobs *entities.Jobs) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.db.InsertOne(ctx, jobs)
	return err
}

func (r *Repository) Get(hash string) (*entities.Jobs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var jobs entities.Jobs
	err := r.db.FindOne(ctx, bson.M{"hash": hash}).Decode(&jobs)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &jobs, err
}

func (r *Repository) Update(jobs *entities.Jobs) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"hash": jobs.Hash}
	update := bson.M{
		"$set": bson.M{
			"positions":  jobs.Positions,
			"updated_at": jobs.UpdatedAt,
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
