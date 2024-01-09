package routes

import (
	"golangsidang/repository"
	"golangsidang/utils"

	"github.com/gofiber/fiber/v2"
)

func ProgramRoutes(app *fiber.App, r *repository.ProgramRepository) {
	api := app.Group("/api")
	api.Use(utils.TokenMiddleware)
	// Apply middleware for token verification to the remaining routes
	api.Post("/createprogram", r.CreateProgram)
}
