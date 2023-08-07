package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/test-netzme/internal/config"
	"github.com/test-netzme/internal/factory/repository"
	mock_user "github.com/test-netzme/internal/mock/repository/api"
	model "github.com/test-netzme/internal/model/entity"
	"github.com/test-netzme/internal/usecase"
)

func TestUser_FindOne_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserAPI := mock_user.NewMockUserAPI(ctrl)
	mockUserAPI.EXPECT().GetRandomUser().Return(model.UserAPIEntity{
		Results: []model.UserAPIEntityResult{
			{
				Gender: "male",
				Name: model.UserAPIEntityName{
					Title: "Mr.",
					First: "Muhammad Cholis",
					Last:  "Malik",
				},
				Location: model.UserAPIEntityLocation{
					Street: model.UserAPIEntityLocationStreet{
						Number: 1,
						Name:   "Jl. perjuangan",
					},
					City: "DKI Jakarta",
				},
				Picture: model.UserAPIEntityPicture{
					Large: "https://github.com/mcholismalik/",
				},
			},
		},
	}, nil)

	cfg := &config.Configuration{}
	mockRepository := repository.Factory{
		UserAPI: mockUserAPI,
	}
	u := usecase.NewUser(cfg, mockRepository)

	result, err := u.FindOne(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUser_FindOne_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserAPI := mock_user.NewMockUserAPI(ctrl)
	mockUserAPI.EXPECT().GetRandomUser().Return(model.UserAPIEntity{}, errors.New("mock err"))

	cfg := &config.Configuration{}
	mockRepository := repository.Factory{
		UserAPI: mockUserAPI,
	}
	u := usecase.NewUser(cfg, mockRepository)

	_, err := u.FindOne(context.Background())

	assert.Error(t, err)
}
