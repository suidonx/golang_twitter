package service

import (
	"context"

	repository "github.com/suidonx/golang_twitter/internal/repository/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	queries *repository.Queries
}

func NewAuthService(q *repository.Queries) *AuthService {
	return &AuthService{queries: q}
}

func (s *AuthService) SignUp(ctx context.Context, email string, password string) error {
	// ハッシュ化
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// repositoryにinsert依頼
	params := repository.CreateUserParams{
		Email:        email,
		PasswordHash: string(passwordHash),
	}
	_, err = s.queries.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	// 結果をControllerに返す
	return nil
}

func Login() {}

// 物理削除
func DeleteAccount() {}
