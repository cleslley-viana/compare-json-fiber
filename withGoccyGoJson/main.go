package main

import (
	"fiber-json-compare/handler"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	port := ":3000"

	app.Post("/", handler.MessageReaderHandler)

	log.Fatal(app.Listen(port))
}
