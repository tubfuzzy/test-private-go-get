package controller

import (
	"fmt"
	"go-circuit-breaker/internal/service"

	"go-circuit-breaker/pkg/circuitbreaker"

	"github.com/gofiber/fiber/v2"
)

func Example(c *fiber.Ctx) error {
	value, err := circuitbreaker.CB.Execute(func() (interface{}, error) {
		if c.Params("id") == "123" {
			fmt.Println(circuitbreaker.CB.Counts().TotalFailures)
			return service.GetB()
		}
		return service.GetA()
	})
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}
	fmt.Println(circuitbreaker.CB.Counts().TotalSuccesses)
	return c.JSON(value)
}

func Example2(c *fiber.Ctx) error {
	value, err := circuitbreaker.CB.Execute(func() (interface{}, error) {
		if c.Params("id") == "123" {
			fmt.Println(circuitbreaker.CB.Counts().TotalFailures)
			return service.GetB()
		}
		return service.GetA()
	})
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}
	fmt.Println(circuitbreaker.CB.Counts().TotalSuccesses)
	return c.JSON(value)
}
