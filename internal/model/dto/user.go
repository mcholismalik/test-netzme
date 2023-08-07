package dto

import (
	model "github.com/test-netzme/internal/model/entity"
	res "github.com/test-netzme/pkg/util/response"
)

// response
type (
	UserResponse struct {
		model.UserEntity
	}
	UserResponseDoc struct {
		Body struct {
			Meta res.Meta     `json:"meta"`
			Data UserResponse `json:"data"`
		} `json:"body"`
	}
)
