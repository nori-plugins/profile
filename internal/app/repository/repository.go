package repository

import (
	"github.com/google/wire"
	"github.com/nori-plugins/profile/internal/repository"
	"github.com/nori-plugins/profile/internal/repository/profile"
)

var RepositorySet = wire.NewSet(
	profile.New,
	repository.New,
)
