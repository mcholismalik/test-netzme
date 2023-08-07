package user

import (
	"encoding/json"
	"net/http"

	"github.com/test-netzme/internal/config"
	model "github.com/test-netzme/internal/model/entity"
)

type (
	UserAPI interface {
		GetRandomUser() (model.UserAPIEntity, error)
	}

	userAPI struct {
		config *config.Configuration
	}
)

func NewUserAPI(cfg *config.Configuration) UserAPI {
	return &userAPI{cfg}
}

func (r *userAPI) GetRandomUser() (results model.UserAPIEntity, err error) {
	url := r.config.Repository.API.User.URL

	response, err := http.Get(url)
	if err != nil {
		return results, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
		return results, err
	}

	return results, nil
}
