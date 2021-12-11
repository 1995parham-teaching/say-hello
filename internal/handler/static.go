package handler

import (
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

const MaxAge = 100

type Static struct {
	FS http.FileSystem
}

func NewStatic(path string, embed fs.FS) Static {
	fsys, err := fs.Sub(embed, path)
	if err != nil {
		panic(err)
	}

	return Static{
		FS: http.FS(fsys),
	}
}

func (s Static) Register(group fiber.Router) {
	group.Use("/", filesystem.New(filesystem.Config{
		Next:         nil,
		Root:         s.FS,
		PathPrefix:   "",
		Browse:       false,
		Index:        "index.html",
		MaxAge:       MaxAge,
		NotFoundFile: "",
	}))
}
