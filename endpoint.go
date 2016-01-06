package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.Redirect("https://twitter.com/_captainpete")
	})

	m.Run()
}
