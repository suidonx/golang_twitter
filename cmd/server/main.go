package main

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/suidonx/golang_twitter/internal/router"
)

func main() {
	r := router.Setup()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validatePassword)
	}

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
