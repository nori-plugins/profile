package service

import (
	"context"

	"github.com/nori-plugins/profile/internal/domain/entity"
)

type ProfileService interface {
	GetByID(ctx context.Context, data GetByIdData) (*entity.Profile, error)
	Create(ctx context.Context, data CreateProfileData) error
}

type GetByIdData struct {
	Id uint64
}

type CreateProfileData struct {
	FirstName string
	LastName  string
	NickName  string
}
