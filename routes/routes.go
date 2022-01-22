package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	now := time.Now()
	fmt.Println(now.Unix())

	app.Get("/", func(c *fiber.Ctx) error {
		timestamp := Return{
			now.Unix(),
		}

		return c.JSON(timestamp)
	})
}

type Return struct {
	Timestamp int64
}
