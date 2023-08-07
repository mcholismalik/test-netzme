package api

import (
	"github.com/gin-gonic/gin"
	"github.com/test-netzme/internal/delivery/api/handler"
	"github.com/test-netzme/internal/factory"
)

func Init(e *gin.Engine, f factory.Factory) {
	// routes
	prefix := "api"

	handler.NewUser(f).Route(e.Group(prefix + "/person"))
}
