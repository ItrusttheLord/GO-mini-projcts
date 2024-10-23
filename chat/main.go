package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Use CORS middleware
	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1879166",
		Key:     "5ba89dc72690fcf46b34",
		Secret:  "9740a9d3f328dd909791",
		Cluster: "us2",
		Secure:  true,
	}

	// Define a route for the POST method on the path '/api/messages'
	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}

		// Trigger the event with Pusher
		err := pusherClient.Trigger("chat", "message", data)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to send message",
			})
		}

		// Send a success response to the client
		return c.JSON(fiber.Map{
			"message": "Message sent successfully",
		})
	})

	// Start the server on port 8000
	log.Fatal(app.Listen(":8000"))
}
