package repository

import "github.com/nori-plugins/profile/internal/domain/repository"

type Repository struct {
	ProfileRepository repository.ProfileRepository
}

type Params struct {
	ProfileRepository repository.ProfileRepository
}

func New(params Params) *Repository {
	repository := Repository{
		ProfileRepository: params.ProfileRepository,
	}
	return &repository
}
