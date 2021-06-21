package service

import (
	"github.com/google/wire"
	"github.com/nori-plugins/profile/internal/service"
	"github.com/nori-plugins/profile/internal/service/profile"
)

var ServiceSet = wire.NewSet(
	wire.Struct(new(profile.Params), "ProfileRepository"),
	profile.New,
	wire.Struct(new(service.Params), "ProfileService"),
	service.New,
)
