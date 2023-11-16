package handler

import (
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MessageRequest struct {
	FirstMessage  string `json:"first_message"`
	SecondMessage string `json:"second_message"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func MessageReaderHandler(c *fiber.Ctx) error {
	req := new(MessageRequest)
	if err := c.BodyParser(req); err != nil {
		slog.Error("falha ao ler dados da request", err)

		c.Status(http.StatusBadRequest).JSON(&ErrorResponse{Error: err.Error()})
	}

	response := &MessageResponse{
		Message: req.FirstMessage + req.SecondMessage,
	}

	return c.JSON(response)
}
