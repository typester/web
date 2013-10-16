package main

import (
	"github.com/typester/web"
	"github.com/typester/web/template"
)

func Root(c *web.Context) {
	template.Render(c, "index.html", map[string]interface{}{})
}

func main() {
	template.SetTemplateDir("template")

	app := web.NewApp()
	app.Handle("/", Root).Methods("POST")
	app.Run(":5000")
}









