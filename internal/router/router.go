package router

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/suidonx/golang_twitter/internal/controller"
)

func Setup(ac *controller.AuthController, tc *controller.TweetController) *gin.Engine {
	router := gin.Default()

	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		log.Fatal("SESSION_SECRETが設定されていません")
	}
	store, err := redis.NewStore(10, "tcp", "redis:6379", "", "", []byte(sessionSecret))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{
		Path:     "/", // サイト全体のリクエストでセッションCookieを使用できる
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
	})
	router.Use(sessions.Sessions("session_id", store)) // Cookie名がsession_id

	router.GET("/health_check", controller.HealthCheck)
	router.POST("/signup", ac.SignUp)
	router.GET("/activate", ac.Activate)
	router.POST("/login", ac.Login)
	router.POST("/post", tc.POST)

	return router
}
