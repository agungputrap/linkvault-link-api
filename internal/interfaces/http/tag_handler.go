package http

import (
	"github.com/agungputrap/linkvault-link-api/internal/application/tag/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TagHandler struct {
	listUserTagUC *usecases.ListUserTagUseCase
}

func NewTagHandler(listUserTagUC *usecases.ListUserTagUseCase) *TagHandler {
	return &TagHandler{listUserTagUC}
}

func (h *TagHandler) ListTags(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	tags, err := h.listUserTagUC.Execute(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(tags)
}
