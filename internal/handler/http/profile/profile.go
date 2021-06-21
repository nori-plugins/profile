package profile

import (
	"net/http"
	"strconv"

	"github.com/nori-plugins/profile/internal/handler/http/response"

	"github.com/go-chi/chi"
	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-plugins/profile/internal/domain/service"
)

type Handler struct {
	profileService service.ProfileService
	logger         logger.FieldLogger
}

type Params struct {
	ProfileService service.ProfileService
	Logger         logger.FieldLogger
}

func New(params Params) *Handler {
	return &Handler{
		profileService: params.ProfileService,
		logger:         params.Logger,
	}
}

func (h *Handler) GetProfileById(w http.ResponseWriter, r *http.Request) {
	/*	_, err := h.cookieHelper.GetSessionID(r)
		if err != nil {
			h.logger.Error("%s", err)
			http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
		}*/

	id := chi.URLParam(r, "id")
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	profile, err := h.profileService.GetByID(r.Context(), service.GetByIdData{Id: u})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, *profile)
}

func (h *Handler) PostProfile(w http.ResponseWriter, r *http.Request) {
	data, err := newCreateProfileData(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = h.profileService.Create(r.Context(), data); err != nil {
		h.logger.Error("%s", err)
		//h.errorHelper.Error(w, err)
	}

	w.WriteHeader(http.StatusCreated)
}
