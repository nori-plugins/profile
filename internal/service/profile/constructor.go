package profile

import (
	"github.com/nori-plugins/profile/internal/domain/repository"
	"github.com/nori-plugins/profile/internal/domain/service"
)

type ProfileService struct {
	profileRepository repository.ProfileRepository
}

type Params struct {
	ProfileRepository repository.ProfileRepository
}

func New(params Params) service.ProfileService {
	return &ProfileService{
		profileRepository: params.ProfileRepository,
	}
}
