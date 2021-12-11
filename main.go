package main

import (
	"embed"
	"log"

	"github.com/cng-by-example/say-hello/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed web/say-hello/build
var embeddedFiles embed.FS

func main() {
	app := fiber.New()

	app.Use(logger.New())

	handler.NewStatic("web/say-hello/build", embeddedFiles).Register(app.Group("/"))
	handler.NewHello().Register(app.Group("/api"))

	log.Fatal(app.Listen(":3000"))
}
