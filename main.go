package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed say-hello/build
var embeddedFiles embed.FS

func main() {
	app := fiber.New()

	app.Use("/", filesystem.New(filesystem.Config{
		Root: getFileSystem(),
	}))
	log.Fatal(app.Listen(":3000"))
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embeddedFiles, "say-hello/build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
