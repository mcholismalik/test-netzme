package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/test-netzme/internal/config"
	api "github.com/test-netzme/internal/delivery/api"
	"github.com/test-netzme/internal/factory"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/test-netzme/docs"
)

func Init(cfg *config.Configuration, f factory.Factory) {
	var (
		APP        = cfg.App.Name
		VERSION    = cfg.App.Version
		DOC_HOST   = cfg.Swagger.SwaggerHost
		DOC_SCHEME = cfg.Swagger.SwaggerScheme
	)

	// engine
	e := gin.Default()

	// index
	e.GET("/", func(c *gin.Context) {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		c.JSON(http.StatusOK, message)
	})

	// doc
	docs.SwaggerInfo.Title = APP
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = DOC_HOST
	docs.SwaggerInfo.Schemes = []string{DOC_SCHEME}
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// delivery
	api.Init(e, f)

	// engine run
	e.Run(":" + cfg.App.Port)
}
