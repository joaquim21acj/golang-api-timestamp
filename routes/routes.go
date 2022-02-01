package routes

import (
	"fmt"
	"main/controllers"
	"main/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	now := time.Now()
	fmt.Println(now.Unix())

	app.Get("/", func(c *fiber.Ctx) error {
		db := database.GetConnection()
		timestamp := controllers.GetTimestamp(db)

		return c.JSON(timestamp)
	})

	app.Get("/requests-success", func(c *fiber.Ctx) error {
		db := database.GetConnection()
		requestsSuccess := controllers.GetAllRequests(db)

		return c.JSON(requestsSuccess)
	})
}

type Return struct {
	Timestamp int64
}
