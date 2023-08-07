package usecase

import (
	"github.com/test-netzme/internal/config"
	"github.com/test-netzme/internal/factory/repository"
	"github.com/test-netzme/internal/usecase"
)

type Factory struct {
	User usecase.User
}

func Init(cfg *config.Configuration, r repository.Factory) Factory {
	f := Factory{}

	f.User = usecase.NewUser(cfg, r)

	return f
}
