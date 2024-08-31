package entities

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Analytics struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Information string             `bson:"information" json:"information"`
	Values      []float64          `bson:"values" json:"values"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewAnalytics(name, information string, values []float64) (*Analytics, error) {
	analytics := &Analytics{
		Name:        name,
		Information: information,
		Values:      values,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := analytics.IsValid(); err != nil {
		return nil, err
	}

	return analytics, nil
}

func (c *Analytics) IsValid() error {
	if c.Name == "" || c.Information == "" || len(c.Values) == 0 {
		return fmt.Errorf("invalid analytics")
	}

	return nil
}
