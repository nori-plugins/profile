package profile

import (
	"context"

	"github.com/nori-plugins/profile/internal/domain/errors"

	"github.com/nori-plugins/profile/internal/domain/entity"

	"github.com/nori-plugins/profile/internal/domain/service"
)

func (srv ProfileService) GetByID(ctx context.Context, data service.GetByIdData) (*entity.Profile, error) {
	/*	if err := data.Validate(); err != nil {
		return nil, err
	}*/

	user, err := srv.profileRepository.FindByID(ctx, data.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.UserNotFound
	}
	return user, nil
}
