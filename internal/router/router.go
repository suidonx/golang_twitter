package router

import (
	"github.com/gin-gonic/gin"
	"github.com/suidonx/golang_twitter/internal/controller"
)

func Setup(ac *controller.AuthController) *gin.Engine {
	router := gin.Default()

	router.GET("/health_check", controller.HealthCheck)
	router.POST("/signup", ac.SignUp)

	return router
}
