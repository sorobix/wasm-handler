package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type request struct {
	Data string `json:"data"`
}

func FileFormatterRest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data request
		err := c.BodyParser(&data)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"success": false,
				"error":   "bad request",
			})
		}
		input, err := base64Decoder(data.Data)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"success": false,
				"error":   "bad base64 input",
			})
		}
		ch := make(chan string)
		go func() {
			runFormatter(string(input), ch)
		}()
		res := <-ch
		if res == "" {
			c.Status(http.StatusNotAcceptable)
			return c.JSON(fiber.Map{
				"success": false,
				"error":   "bad rust code",
				"logs":    err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"success": true,
			"data":    base64Encoder(string(res)),
		})
	}
}
