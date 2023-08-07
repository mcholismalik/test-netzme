package usecase

import (
	"context"

	"github.com/test-netzme/internal/config"

	"github.com/test-netzme/internal/factory/repository"
	"github.com/test-netzme/internal/model/dto"
	res "github.com/test-netzme/pkg/util/response"
)

type User interface {
	FindOne(ctx context.Context) (dto.UserResponse, error)
}

type user struct {
	Cfg  *config.Configuration
	Repo repository.Factory
}

func NewUser(cfg *config.Configuration, f repository.Factory) User {
	return &user{cfg, f}
}

func (u *user) FindOne(ctx context.Context) (result dto.UserResponse, err error) {
	randomUsers, err := u.Repo.UserAPI.GetRandomUser()
	if err != nil {
		return result, err
	}

	if len(randomUsers.Results) == 0 {
		return result, res.ErrorBuilder(res.Constant.Error.NotFound, err)
	}

	result.UserEntity = randomUsers.ToUserEntity()

	return result, nil
}
