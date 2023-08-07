package factory

import (
	"github.com/test-netzme/internal/config"
	"github.com/test-netzme/internal/factory/repository"
	"github.com/test-netzme/internal/factory/usecase"
)

type Factory struct {
	Repository repository.Factory
	Usecase    usecase.Factory
}

func Init(cfg *config.Configuration) (Factory, error) {
	f := Factory{}

	// repository
	f.Repository = repository.Init(cfg)

	// usecase
	f.Usecase = usecase.Init(cfg, f.Repository)

	return f, nil
}
