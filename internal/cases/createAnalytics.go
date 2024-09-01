package cases

import (
	"github.com/renan-parise/xcal-analytics/internal/analytes"
	"github.com/renan-parise/xcal-analytics/internal/entities"
	"github.com/renan-parise/xcal-analytics/internal/repositories"
)

type createAnalyticsCase struct {
	repository repositories.IAnalyticsRepository
}

func NewCreateAnalyticsCase(repository repositories.IAnalyticsRepository) *createAnalyticsCase {
	return &createAnalyticsCase{repository: repository}
}

func (u *createAnalyticsCase) Execute(name string, analytes analytes.Analytes) error {
	analytics, err := entities.NewAnalytics(name, analytes)
	if err != nil {
		return err
	}

	if err := analytics.Check(); err != nil {
		return err
	}

	err = u.repository.Insert(analytics)
	if err != nil {
		return err
	}

	return nil
}
