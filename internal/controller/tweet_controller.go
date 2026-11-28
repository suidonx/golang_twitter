package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suidonx/golang_twitter/internal/service"
)

type TweetController struct {
	tweetService *service.TweetService
}

func NewTweetController(t *service.TweetService) *TweetController {
	return &TweetController{tweetService: t}
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

	userID, ok := session.Get("user_id").(uuid.UUID)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}

	err := tc.tweetService.Post(c.Request.Context(), req.Content, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "content": req.Content})
	return
}

func ListTweets(c *gin.Context) {}

func GetTweet(c *gin.Context) {}

func LikeTweet(c *gin.Context) {}

func Retweet(c *gin.Context) {}

func FollowUser(c *gin.Context) {}

func BookmarkTweet(c *gin.Context) {}
