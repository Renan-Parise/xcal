package repositories

import "github.com/renan-parise/gofreela/internal/entities"

type IJobsRepository interface {
	Insert(*entities.Jobs) error
	Get(hash string) (*entities.Jobs, error)
	Update(*entities.Jobs) error
}
