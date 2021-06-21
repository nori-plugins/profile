package profile

import (
	"encoding/json"
	"net/http"

	"github.com/nori-plugins/profile/internal/domain/service"
)

type CreateProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	NickName  string `json:"nick_name"`
}

func newCreateProfileData(r *http.Request) (service.CreateProfileData, error) {
	var body CreateProfileRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return service.CreateProfileData{}, err
	}
	return service.CreateProfileData{}, nil
}
