package service

import (
	"context"
	"database/sql"

	repository "github.com/suidonx/golang_twitter/internal/repository/sqlc"
)

type TweetService struct {
	queries *repository.Queries
	db      *sql.DB
}

func (ts *TweetService) Post(ctx context.Context, email string, password string) error {
}
