package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cng-by-example/say-hello/internal/response"
	"github.com/gofiber/fiber/v2"
)

type Hello struct{}

func NewHello() Hello {
	return Hello{}
}

// nolint: wrapcheck
func (h Hello) Get(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(response.Message{
		Msg: fmt.Sprintf("Hello at %s", time.Now().Format(time.RFC3339)),
	})
}

func (h Hello) Register(group fiber.Router) {
	group.Get("/hello", h.Get)
}
