package entities

import (
	"fmt"
	"time"

	"github.com/renan-parise/gofreela/internal/analytes"
	"github.com/renan-parise/gofreela/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Analytics struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Hash      string             `bson:"hash" json:"hash"`
	Analytes  analytes.Analytes  `bson:"analytes" json:"analytes"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewAnalytics(hash string, analytes analytes.Analytes) (*Analytics, error) {
	currentTime := *utils.Now()

	analytes.PopulateIncludedAtCreation(currentTime)

	analytics := &Analytics{
		Hash:      hash,
		Analytes:  analytes,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if err := analytics.Check(); err != nil {
		return nil, err
	}

	return analytics, nil
}

func (c *Analytics) Check() error {
	if c.Hash == "" {
		return fmt.Errorf("invalid hash instance")
	}

	return nil
}
