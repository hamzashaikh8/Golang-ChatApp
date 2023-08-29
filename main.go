package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	app := fiber.New()

	app.Use(cors.New()) //Provide middleware to app to allow CORS requests from any origin to any route

	// Pusher client for real-time communication between server and client
	pusherClient := pusher.Client{
		AppID:   "1659280",
		Key:     "fadf38e1e146d8dfd05d",
		Secret:  "08f28e9506c223a07900",
		Cluster: "us2",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		pusherClient.Trigger("chat", "message", data)
		return c.JSON([]string{})
	})

	app.Listen(":8000")
}
