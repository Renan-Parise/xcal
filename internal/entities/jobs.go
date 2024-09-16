package entities

import (
	"fmt"
	"time"

	"github.com/renan-parise/gofreela/internal/positions"
	"github.com/renan-parise/gofreela/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Jobs struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Hash      string              `bson:"hash" json:"hash"`
	Positions positions.Positions `bson:"positions" json:"positions"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at" json:"updated_at"`
}

func NewJobs(hash string, positions positions.Positions) (*Jobs, error) {
	currentTime := *utils.Now()

	positions.PopulateIncludedAtCreation(currentTime)

	jobs := &Jobs{
		Hash:      hash,
		Positions: positions,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if err := jobs.Check(); err != nil {
		return nil, err
	}

	return jobs, nil
}

func (c *Jobs) Check() error {
	if c.Hash == "" {
		return fmt.Errorf("invalid hash instance")
	}

	return nil
}
