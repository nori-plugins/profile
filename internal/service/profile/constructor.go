package profile

import (
	"github.com/nori-plugins/profile/internal/domain/service"
)

type ProfileService struct {
}

type Params struct {
}

func New(params Params) service.ProfileService {
	return &ProfileService{}
}
