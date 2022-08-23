package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func NewSessionize(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		sess.Set("name", c.Query("name", "unknown user"))
		if err := sess.Save(); err != nil {
			panic(err)
		}
		return c.Next()
	}
}
