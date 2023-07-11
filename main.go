package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	InitialiseRedis()
	log.Println("Redis has been Initialised")
	app := fiber.New()
	app.Use(cors.New())

	handler := app.Group("/handler")
	handler.Get("/heartbeat", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":           "Running",
			"redis_connection": RedisConnected(),
		})
	})
	handler.Get("/format", FileFormatterRest())

	ws := handler.Group("/ws")
	ws.Use("/", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	ws.Get("/compile", websocket.New(compilerController()))

	app.Listen(":" + wsPort)
	log.Println("App is listening on Port:", wsPort)

}
