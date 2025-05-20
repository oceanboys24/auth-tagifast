package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/oceanboys24/auth-tagifast/configs"
	"github.com/oceanboys24/auth-tagifast/database"
	auth "github.com/oceanboys24/auth-tagifast/internal/auth/handler"
	"github.com/oceanboys24/auth-tagifast/internal/auth/repository"
	"github.com/oceanboys24/auth-tagifast/internal/auth/routes"
	"github.com/oceanboys24/auth-tagifast/internal/auth/service"
)

func main() {
	cfg, err := configs.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	gormdb, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	db, err := gormdb.DB()
	if err != nil {
		log.Fatalf("failed to initialize DB: %v", err)
	}

	defer db.Close()

	app := fiber.New()

	authRepo := repository.NewAuthRepository(gormdb)
	authService := service.NewAuthService(authRepo)
	AuthHandler := auth.NewAuthHandler(*authService)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Welcome to Auth-Supabase!",
		})
	})

	apiRoute := app.Group("/v1")

	routes.AuthRoutes(apiRoute.Group("/auth"), AuthHandler)

	log.Println("🚀 Server running at http://localhost:4000")
	app.Listen(":4000")
}
