package app

import (
	"github.com/google/wire"
	"github.com/nori-plugins/profile/internal/app/handler"
	"github.com/nori-plugins/profile/internal/app/repository"
	"github.com/nori-plugins/profile/internal/app/service"
	"github.com/nori-plugins/profile/internal/app/transactor"
)

var AppSet = wire.NewSet(transactor.TransactorSet, handler.HandlerSet, repository.RepositorySet, service.ServiceSet)
