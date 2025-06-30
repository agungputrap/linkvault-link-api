package http

import (
	"github.com/agungputrap/linkvault-link-api/internal/application/link/dto"
	"github.com/agungputrap/linkvault-link-api/internal/application/link/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

type LinkHandler struct {
	createUC *usecases.CreateLinkUseCase
	getUC    *usecases.GetLinksUseCase
	deleteUC *usecases.DeleteLinksUseCase
	updateUC *usecases.UpdateLinksUseCase
}

func NewLinkHandler(createUC *usecases.CreateLinkUseCase, getUC *usecases.GetLinksUseCase, deleteUC *usecases.DeleteLinksUseCase, updateUC *usecases.UpdateLinksUseCase) *LinkHandler {
	return &LinkHandler{createUC, getUC, deleteUC, updateUC}
}

func (h *LinkHandler) Create(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var req dto.CreateLinkRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	res, err := h.createUC.Execute(c.Context(), userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *LinkHandler) GetAll(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	links, err := h.getUC.Execute(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(links)
}

func (h *LinkHandler) Delete(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	linkID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid link ID",
		})
	}
	if err := h.deleteUC.Execute(c.Context(), uint(linkID), userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *LinkHandler) Update(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	linkID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid link ID",
		})
	}

	var req dto.UpdateLinkRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body request",
		})
	}

	res, err := h.updateUC.Execute(c.Context(), uint(linkID), userId, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(res)
}
