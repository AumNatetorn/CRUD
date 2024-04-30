package app

import (
	"CRUD/configs"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	HealthPath = "/health"
)

type server struct {
	App *fiber.App
}

func NewServer(cfg configs.Config, pingFn ...func() error) *server {
	f := fiber.New(
		fiber.Config{
			ReadTimeout:           5 * time.Second,
			WriteTimeout:          5 * time.Second,
			IdleTimeout:           30 * time.Second,
			DisableStartupMessage: true,
			CaseSensitive:         true,
			StrictRouting:         true,
		},
	)

	f.Use(HealthPath, healthDependencies(pingFn...))

	return &server{f}
}

func (s *server) Start(port int) error {
	err := s.App.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return nil
}

func healthDependencies(pingFn ...func() error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, fn := range pingFn {
			err := fn()
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"status": "internal server error",
				})
			}
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status": "healthy",
		})
	}
}
