package handler

import (
	"github.com/google/wire"
	httpHandler "github.com/nori-plugins/profile/internal/handler/http"
	"github.com/nori-plugins/profile/internal/handler/http/profile"
)

var HandlerSet = wire.NewSet(
	wire.Struct(new(httpHandler.Handler), "R", "ProfileHandler"),
	wire.Struct(new(profile.Params), "ProfileService", "Logger"),
	profile.New,
)
