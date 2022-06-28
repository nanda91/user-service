package usecase

import (
	"user-service/config/env"
	models "user-service/model"
	"user-service/repository"
)

type UseCaseInterface interface {
	FindAllUser()([]models.User, error)
	FindUser(userId string) (models.User, error)
}

type UseCase struct {
	Config env.Config
	Repo repository.RepositoryInterface
}


func NewUseCase(config env.Config, repo repository.RepositoryInterface) UseCaseInterface {
	return &UseCase{
		Config: config,
		Repo: repo,
	}
}

func (uc *UseCase) FindAllUser() ([]models.User, error) {
	var (
		data []models.User
		err error
	)

	data, err = uc.Repo.FindAllUser()
	return data, err
}

func (uc *UseCase) FindUser(userId string) (models.User, error) {
	var (
		data models.User
		err error
	)

	data, err = uc.Repo.FindUser(userId)
	return data, err
}