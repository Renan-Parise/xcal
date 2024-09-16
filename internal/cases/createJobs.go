package cases

import (
	"github.com/renan-parise/gofreela/internal/entities"
	"github.com/renan-parise/gofreela/internal/positions"
	"github.com/renan-parise/gofreela/internal/repositories"
)

type createJobsCase struct {
	repository repositories.IJobsRepository
}

func NewCreateJobsCase(repository repositories.IJobsRepository) *createJobsCase {
	return &createJobsCase{repository: repository}
}

func (u *createJobsCase) Execute(name string, positions positions.Position) error {
	jobs, err := entities.NewJobs(name, positions)
	if err != nil {
		return err
	}

	if err := jobs.Check(); err != nil {
		return err
	}

	err = u.repository.Insert(jobs)
	if err != nil {
		return err
	}

	return nil
}
