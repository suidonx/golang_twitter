package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/suidonx/golang_twitter/internal/service"
)

type TweetController struct {
	tweetService *service.TweetService
}

type PostRequest struct {
	Content string `json:"content" binding:"required,max=140,min=1"`
}

func (tc *TweetController) POST(c *gin.Context) {
	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)

	userID := session.Get("user_id")
	if userID != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
	}

	req.Content

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "content": req.Content})
	return
}

func ListTweets(c *gin.Context) {}

func GetTweet(c *gin.Context) {}

func LikeTweet(c *gin.Context) {}

func Retweet(c *gin.Context) {}

func FollowUser(c *gin.Context) {}

func BookmarkTweet(c *gin.Context) {}
