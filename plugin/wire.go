//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-io/common/v5/pkg/domain/registry"
	noriGorm "github.com/nori-io/interfaces/database/gorm"
	noriHttp "github.com/nori-io/interfaces/nori/http"
	"github.com/nori-plugins/profile/internal/app"
	config2 "github.com/nori-plugins/profile/internal/config"
	httpHandler "github.com/nori-plugins/profile/internal/handler/http"
)

var set = wire.NewSet(
	noriGorm.GetGorm,
	noriHttp.GetRouter)

func Initialize(registry registry.Registry, config config2.Config, logger logger.FieldLogger) (*httpHandler.Handler, error) {
	wire.Build(app.AppSet, set)
	return &httpHandler.Handler{}, nil
}
