package http

import (
	"github.com/agungputrap/linkvault-link-api/internal/application/user/dto"
	"github.com/agungputrap/linkvault-link-api/internal/application/user/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type UserHandler struct {
	registerUC *usecases.RegisterUseCase
	loginUC    *usecases.LoginUseCase
}

func NewUserHandler(registerUC *usecases.RegisterUseCase, loginUC *usecases.LoginUseCase) *UserHandler {
	return &UserHandler{registerUC, loginUC}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	res, err := h.registerUC.Execute(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request"})
	}
	res, err := h.loginUC.Execute(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// TODO: add JWT token here
	claims := jwt.MapClaims{
		"user_id": res.ID,
		"email":   res.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": signedToken,
		"user":  res,
	})
}
