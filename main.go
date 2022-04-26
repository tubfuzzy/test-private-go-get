package main

import (
	"go-circuit-breaker/internal/controller"
	"go-circuit-breaker/pkg/circuitbreaker"

	"github.com/gofiber/fiber/v2"
)

// git config --global url."https://theeraprasert:z4bVgp7jqh7tTsUL5CaP@bitbucket.org".insteadOf "https://bitbucket.org/"
// export BITBUCKET_USERNAME=theeraprasert
// export BITBUCKET_APP_PASSWORD=m8HyRz6AqJYWr8Q73DGk
// git config --global url."https://${BITBUCKET_USERNAME}:${BITBUCKET_APP_PASSWORD}@bitbucket.org".insteadOf "https://bitbucket.org/"
// export GOPRIVATE=bitbucket.org/panda-development
// git config --global url."https://theeraprasert@bitbucket.com:m8HyRz6AqJYWr8Q73DGk@bitbucket.com/".insteadOf "https://bitbucket.com/"

func main() {
	app := fiber.New(fiber.Config{})
	circuitbreaker.Init()
	app.Get("/example/:id", controller.Example)
	app.Get("/example3/:id", controller.Example2)

	if appErr := app.Listen(":" + "8080"); appErr != nil {
		panic(appErr)
	}
}
