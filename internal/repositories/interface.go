package repositories

import "github.com/renan-parise/xcal-analytics/internal/entities"

type IAnalyticsRepository interface {
	Insert(*entities.Analytics) error
	Get(hash string) (*entities.Analytics, error)
	Update(*entities.Analytics) error
}
