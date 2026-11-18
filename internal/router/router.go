package router

import (
	"github.com/gin-gonic/gin"
	"github.com/suidonx/golang_twitter/internal/controller"
)

func Setup() *gin.Engine {
	router := gin.Default()

	router.GET("/health_check", controller.HealthCheck)

	return router
}
