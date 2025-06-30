package http

import (
	linkUC "github.com/agungputrap/linkvault-link-api/internal/application/link/usecases"
	tagUC "github.com/agungputrap/linkvault-link-api/internal/application/tag/usecases"
	userUC "github.com/agungputrap/linkvault-link-api/internal/application/user/usecases"
	"github.com/agungputrap/linkvault-link-api/internal/infrastructure/postgres"
	"github.com/agungputrap/linkvault-link-api/internal/interfaces/http/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := postgres.NewUserRepository(db)
	registerUC := userUC.NewRegisterUseCase(userRepo)
	loginUC := userUC.NewLoginUseCase(userRepo)
	userHandler := NewUserHandler(registerUC, loginUC)

	auth := app.Group("/auth")
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	linkRepo := postgres.NewLinkRepository(db)
	createLinkUC := linkUC.NewCreateLinkUseCase(linkRepo)
	getLinkUC := linkUC.NewGetLinksUseCase(linkRepo)
	deleteLinkUC := linkUC.NewDeleteLinksUseCase(linkRepo)
	updateLinkUC := linkUC.NewUpdateLinksUseCase(linkRepo)
	linkHandler := NewLinkHandler(createLinkUC, getLinkUC, deleteLinkUC, updateLinkUC)

	tagRepo := postgres.NewTagRepository(db)
	listUserTagUC := tagUC.NewListUserTagUseCase(tagRepo)
	tagHandler := NewTagHandler(listUserTagUC)

	api := app.Group("/api", middleware.Protected())
	api.Post("links", linkHandler.Create)
	api.Get("links", linkHandler.GetAll)
	api.Delete("links/:id", linkHandler.Delete)
	api.Put("links/:id", linkHandler.Update)

	api.Get("/tags", tagHandler.ListTags)
}
