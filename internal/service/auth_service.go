package service

import (
	"golang.org/x/crypto/bcrypt"
)

func SignUp(email string, password string) error {
	// ハッシュ化
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// repositoryにinsert依頼

	// 失敗時にerror

	// 結果をControllerに返す
	return err
}

func Login() {}

// 物理削除
func DeleteAccount() {}
