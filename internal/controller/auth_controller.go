package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/suidonx/golang_twitter/internal/service"
)

type SignUpRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required,password"`
}

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(a *service.AuthService) *AuthController {
	return &AuthController{authService: a}
}

func (ac *AuthController) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ac.authService.SignUp(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "OK"})
	return
}

func (ac *AuthController) Activate(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "認証トークンが設定されていません"})
		return
	}

	err := ac.authService.Activate(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "mail activasion successed"})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)

	userID, err := ac.authService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "メールアドレスまたはパスワードが正しくありません"})
		return
	}

	session.Set("user_id", userID.String())
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "セッションの保存に失敗しました",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ログインに成功しました"})
}

// 物理削除
func DeleteAccount(c *gin.Context) {}
