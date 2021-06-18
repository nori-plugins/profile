package service

import (
	"context"

	"github.com/nori-plugins/profile/internal/domain/entity"
)

type ProfileService interface {
	GetByID(ctx context.Context, data GetByIdData) (*entity.Profile, error)
}

type GetByIdData struct {
	Id uint64
}
