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

func SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.SignUp(req.Email, req.Password)
	c.JSON(http.StatusCreated, gin.H{"status": "OK"})

}

func Login(c *gin.Context) {}

// 物理削除
func DeleteAccount(c *gin.Context) {}
