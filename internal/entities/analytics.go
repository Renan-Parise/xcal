package entities

import (
	"fmt"
	"time"

	"github.com/renan-parise/xcal-analytics/internal/analytes"
	"github.com/renan-parise/xcal-analytics/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Analytics struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User      string             `bson:"user" json:"user"`
	Hash      string             `bson:"hash" json:"hash"`
	Analytes  analytes.Analytes  `bson:"analytes" json:"analytes"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewAnalytics(user, hash string, analytes analytes.Analytes) (*Analytics, error) {
	currentTime := *utils.Now()

	utils.PopulateIncludedAtCreation(&analytes, currentTime)

	analytics := &Analytics{
		User:      user,
		Hash:      hash,
		Analytes:  analytes,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if err := analytics.IsValid(); err != nil {
		return nil, err
	}

	return analytics, nil
}

func (c *Analytics) IsValid() error {
	if c.User == "" || c.Hash == "" {
		return fmt.Errorf("invalid analytics instance")
	}

	return nil
}
