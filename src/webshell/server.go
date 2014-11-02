package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "view",
		Extensions: []string{".html"},
	}))
	m.Get("/", func(r render.Render) {
		r.HTML(200, "shell", "shell")
	})

	m.Run()
}
