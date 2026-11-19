package controller

import (
	"net/http"

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

func Login(c *gin.Context) {}

// 物理削除
func DeleteAccount(c *gin.Context) {}
