package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/test-netzme/internal/factory"
	res "github.com/test-netzme/pkg/util/response"
)

type (
	user struct {
		Factory factory.Factory
	}
	User interface {
		Route(g *gin.RouterGroup)
		GetOne(c *gin.Context)
	}
)

func NewUser(f factory.Factory) User {
	return &user{f}
}

func (h *user) Route(g *gin.RouterGroup) {
	g.GET("", h.GetOne)
}

// Get user
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /api/person [get]
func (h *user) GetOne(c *gin.Context) {
	result, err := h.Factory.Usecase.User.FindOne(c.Request.Context())
	if err != nil {
		res.ErrorResponse(err).Send(c)
	}

	res.CustomSuccessBuilder(200, result, "Get person success", nil).Send(c)
}
