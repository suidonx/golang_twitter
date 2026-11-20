package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/suidonx/golang_twitter/config"
	"github.com/suidonx/golang_twitter/internal/controller"
	repository "github.com/suidonx/golang_twitter/internal/repository/sqlc"
	"github.com/suidonx/golang_twitter/internal/router"
	"github.com/suidonx/golang_twitter/internal/service"
)

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("password", validatePassword); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("unexpected error")
	}

	cfg := config.Load()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err) // エラーを記録して終了する
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := repository.New(db)
	authService := service.NewAuthService(queries, db)
	authController := controller.NewAuthController(authService)
	r := router.Setup(authController)

	r.Run(":8080")
}

var (
	upperPattern  = regexp.MustCompile(`[A-Z]`)
	lowerPattern  = regexp.MustCompile(`[a-z]`)
	digitPattern  = regexp.MustCompile(`[0-9]`)
	symbolPattern = regexp.MustCompile(`[!?_-]`)
	lengthPattern = regexp.MustCompile(`^.{8,}$`)
)

var validatePassword validator.Func = func(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if ok {
		if !upperPattern.MatchString(password) {
			return false
		}
		if !lowerPattern.MatchString(password) {
			return false
		}
		if !digitPattern.MatchString(password) {
			return false
		}
		if !symbolPattern.MatchString(password) {
			return false
		}
		if !lengthPattern.MatchString(password) {
			return false
		}
		return true
	}
	return false
}
