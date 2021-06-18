package http

import (
	"github.com/nori-plugins/profile/internal/handler/http/profile"

	"github.com/nori-io/interfaces/nori/http"
)

type Handler struct {
	R              http.Http
	ProfileHandler *profile.Handler
}

type Params struct {
	R              http.Http
	ProfileHandler *profile.Handler
}

func New(params Params) *Handler {
	handler := Handler{
		R:              params.R,
		ProfileHandler: params.ProfileHandler,
	}

	// todo: add middleware
	handler.R.Get("/profile.go/{id}", handler.ProfileHandler.GetProfileById)
	//handler.R.Post("/auth/signin", handler.ProfileHandler.PostProfile)
	return &handler
}
