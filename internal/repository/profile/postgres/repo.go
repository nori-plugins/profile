package postgres

import (
	"context"

	"gorm.io/gorm"

	"github.com/nori-plugins/profile/pkg/errors"

	"github.com/nori-plugins/profile/internal/domain/entity"
	"github.com/nori-plugins/profile/internal/transactor"
)

type ProfileRepository struct {
	Tx transactor.Transactor
}

func (r *ProfileRepository) Create(ctx context.Context, e *entity.Profile) error {
	m := newModel(e)

	if err := r.Tx.GetDB(ctx).Create(m).Error; err != nil {
		return errors.NewInternal(err)
	}

	*e = *m.convert()

	return nil
}

func (r *ProfileRepository) FindByID(ctx context.Context, id uint64) (*entity.Profile, error) {
	out := &model{}
	err := r.Tx.GetDB(ctx).Where("id=?", id).First(out).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}
