package http

import (
	"io"

	"github.com/nori-plugins/profile/internal/handler/http/profile"

	netHttp "net/http"

	"github.com/nori-io/interfaces/nori/http"
)

type Handler struct {
	R              http.Router
	ProfileHandler *profile.Handler
}

type Params struct {
	R              http.Router
	ProfileHandler *profile.Handler
}

func New(params Params) *Handler {
	handler := Handler{
		R:              params.R,
		ProfileHandler: params.ProfileHandler,
	}

	handler.R.Route("/profiles", func(r http.Router) {
		r.Post("/", handler.ProfileHandler.PostProfile) // POST /profiles
		// Subrouters:
		r.Route("/{id}", func(r http.Router) {
			r.Use(StubCtx)
			r.Get("/", handler.ProfileHandler.GetProfileById) // GET /profiles/123
		})
	})
	return &handler
}

func StubCtx(next netHttp.Handler) netHttp.Handler {
	return netHttp.HandlerFunc(func(w netHttp.ResponseWriter, r *netHttp.Request) {
		io.WriteString(w, "Use")
		next.ServeHTTP(w, r)
	})
}
