package routes

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/oceanboys24/auth-tagifast/internal/auth/handler"
	"github.com/oceanboys24/auth-tagifast/middleware"
)

func AuthRoutes(app fiber.Router, handler *auth.AuthHandler) {
	app.Post("/v1/auth/login", handler.LoginHandler)
	app.Post("/v1/auth/register", handler.RegisterHandler)

	// check auth
	app.Get("/v1/auth/check", middleware.Auth, handler.AuthCheck)
}
