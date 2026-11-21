package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/smtp"
	"net/url"
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
		TokenHash: hex.EncodeToString(hashedActivationToken[:]),
		ExpiresAt: time.Now().Add(time.Minute * 10),
	}

	err = qtx.CreateToken(ctx, tokenParams)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	u, err := url.Parse("http://localhost:8080/activate")
	if err != nil {
		return err
	}
	q := u.Query()
	q.Set("token", hex.EncodeToString(activationToken))
	u.RawQuery = q.Encode()
	activationURL := u.String()

	err = SendActivationEmail(email, activationURL) // /activate?token=abc123
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Activate(ctx context.Context, token string) error {
	decodeToken, err := hex.DecodeString(token)
	if err != nil {
		return err
	}
	decodeTokenToHash := sha256.Sum256(decodeToken)
	tokenHash := hex.EncodeToString(decodeTokenToHash[:])

	activation, err := s.queries.FindEmailActivationByTokenHash(ctx, tokenHash)
	if err != nil {
		return err
	}

	if time.Now().After(activation.ExpiresAt) {
		return errors.New("認証リンクの有効期限切れです")
	}

	err = s.queries.MarkEmailAsVerified(ctx, activation.UserID)
	if err != nil {
		return err
	}

	return nil
}

func SendActivationEmail(to string, activationURL string) error {
	// メール送信
	from := "no-reply@example.com"
	subject := "Verify your email address"
	body := fmt.Sprintf(`Twitter Cloneにご登録いただきありがとうございます。
メールアドレスの認証を完了するため、以下のリンクを開いてください。

%s

このリンクの有効期限は10分間です。
このメールにお心当たりがない場合は、破棄してください。

Twitter Clone`, activationURL)
	msg := fmt.Sprintf(
		"To: %s\r\n"+
			"From: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/plain; charset=UTF-8\r\n"+
			"\r\n"+
			"%s",
		to,
		from,
		subject,
		body,
	)

	addr := "mailpit:1025"
	err := smtp.SendMail(addr, nil, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func Login() {}

// 物理削除
func DeleteAccount() {}
