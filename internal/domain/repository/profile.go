package repository

import (
	"context"

	"github.com/nori-plugins/profile/internal/domain/entity"
)

type ProfileRepository interface {
	Create(ctx context.Context, e *entity.Profile) error
	FindByID(ctx context.Context, ID uint64) (*entity.Profile, error)
}
