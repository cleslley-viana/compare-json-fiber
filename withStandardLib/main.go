package main

import (
	"fiber-json-compare/handler"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := ":3000"

	app.Post("/", handler.MessageReaderHandler)

	log.Fatal(app.Listen(port))
}
