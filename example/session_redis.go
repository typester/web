package main

import (
	"fmt"
	"github.com/typester/web"
	"github.com/typester/web/session"
	"github.com/typester/web/session/cookie"
	"github.com/typester/web/session/redis"
	"time"
)

func Counter(c *web.Context) {
	session, _ := session.Restore(c)

	var count int64
	data := session.Get("counter")

	if c, ok := data.(int64); ok {
		count = c
	} else {
		count = 0
	}

	count++

	session.Set("counter", count)
	session.Save(c)

	c.Write([]byte(fmt.Sprintf("count: %d", count)))
}

func main() {
	state := &cookie.CookieState{Name: "session", HttpOnly: true}
	store := redis.NewRedisStore(":6379", "session:")

	session.Setup(state, store, "secret", time.Hour*24*365)

	app := web.NewApp()
	app.Handle("/", Counter)
	app.Run(":5000")
}
