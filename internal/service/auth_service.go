package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"time"

	repository "github.com/suidonx/golang_twitter/internal/repository/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	queries *repository.Queries
	db      *sql.DB
}

func NewAuthService(q *repository.Queries, db *sql.DB) *AuthService {
	return &AuthService{queries: q, db: db}
}

func (s *AuthService) SignUp(ctx context.Context, email string, password string) error {
	// ハッシュ化
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userParams := repository.CreateUserParams{
		Email:        email,
		PasswordHash: string(passwordHash),
	}

	// 認証トークン生成
	activationToken := make([]byte, 32)
	rand.Read(activationToken)
	hashedActivationToken := sha256.Sum256(activationToken)

	// repositoryにinsert依頼
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := s.queries.WithTx(tx)

	userId, err := qtx.CreateUser(ctx, userParams)
	if err != nil {
		return err
	}

	tokenParams := repository.CreateTokenParams{
		UserID:    userId,
		TokenHash: hex.EncodeToString(hashedActivationToken),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	err = qtx.CreateToken(ctx, tokenParams)
	if err != nil {
		return err
	}
	return tx.Commit() // 結果をControllerに返す

}

func Login() {}

// 物理削除
func DeleteAccount() {}
