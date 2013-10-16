package main

import (
	"github.com/typester/web"
)

func RootHandler(c *web.Context) {
	c.Write([]byte("Hello"))
}

func main() {
	app := web.NewApp()
	app.Handle("/", RootHandler)
	app.Run(":5000")
}
