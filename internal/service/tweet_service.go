package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	repository "github.com/suidonx/golang_twitter/internal/repository/sqlc"
)

type TweetService struct {
	queries *repository.Queries
	db      *sql.DB
}

func NewTweetService(q *repository.Queries, db *sql.DB) *TweetService {
	return &TweetService{queries: q, db: db}
}

func (ts *TweetService) Post(ctx context.Context, content string, userID uuid.UUID) error {
	postParmas := repository.CreatePostParams{
		Content: content,
		UserID:  userID,
	}

	err := ts.queries.CreatePost(ctx, postParmas)
	if err != nil {
		return err
	}

	return nil
}
