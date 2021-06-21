package service

import (
	"github.com/nori-plugins/profile/internal/domain/service"
)

type Service struct {
	ProfileService service.ProfileService
}

type Params struct {
	ProfileService service.ProfileService
}

func New(params Params) *Service {
	return &Service{
		ProfileService: params.ProfileService,
	}
}
