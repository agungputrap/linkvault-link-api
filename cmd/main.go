package main

import (
	"github.com/agungputrap/linkvault-link-api/internal/infrastructure/database"
	routes "github.com/agungputrap/linkvault-link-api/internal/interfaces/http"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := database.InitPostgres()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	app := fiber.New()
	routes.SetupRoutes(app, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
