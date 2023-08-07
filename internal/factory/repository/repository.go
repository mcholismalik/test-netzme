package repository

import (
	"github.com/test-netzme/internal/config"
	apiRepository "github.com/test-netzme/internal/repository/api"
)

type Factory struct {
	UserAPI apiRepository.UserAPI
}

func Init(cfg *config.Configuration) Factory {
	f := Factory{}

	f.UserAPI = apiRepository.NewUserAPI(cfg)

	return f
}
