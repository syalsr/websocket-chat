package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Payload struct {
	Message string `json:"message"`
	Action  string `json:"action"`
}

func main() {
	app := fiber.New()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		in := Payload{Action: "act", Message: "d"}
		err := c.WriteJSON(in)
		if err != nil {
			log.Print(err)
		}
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	log.Fatal(app.Listen(":8080"))
}
