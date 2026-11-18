package main

import (
	"github.com/suidonx/golang_twitter/internal/router"
)

func main() {
	r := router.Setup()
	r.Run(":8080")
}
