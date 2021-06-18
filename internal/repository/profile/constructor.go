package profile

import (
	"github.com/nori-plugins/profile/internal/domain/repository"
	"github.com/nori-plugins/profile/internal/repository/profile/postgres"
	"github.com/nori-plugins/profile/internal/transactor"
)

func New(tx transactor.Transactor) repository.ProfileRepository {
	return &postgres.ProfileRepository{
		Tx: tx,
	}
}
