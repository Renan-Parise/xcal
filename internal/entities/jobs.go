package entities

import (
	"time"

	"github.com/renan-parise/gofreela/errors"
	"github.com/renan-parise/gofreela/internal/positions"
	"github.com/renan-parise/gofreela/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Jobs struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Hash      string             `bson:"hash" json:"hash"`
	Positions positions.Position `bson:"positions" json:"positions"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewJobs(hash string, positions positions.Position) (*Jobs, error) {
	currentTime := *utils.Now()

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
	if c.Positions.Salary == 0 || c.Positions.Currency == "" {
		return errors.NewValidationError("salary and currency", "salary and currency are required")
	}

	if c.Positions.Title == "" {
		return errors.NewValidationError("title", "title is required")
	}

	if c.Positions.Company == "" {
		return errors.NewValidationError("company", "company is required")
	}

	if c.Positions.Location == "" {
		return errors.NewValidationError("location", "location is required")
	}

	if c.Positions.Description == "" {
		return errors.NewValidationError("description", "description is required")
	}

	return nil
}
