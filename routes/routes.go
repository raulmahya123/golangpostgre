// routes/routes.go
package routes

import (
	"golangsidang/repository"
	"golangsidang/utils"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, r *repository.Repository) {
	api := app.Group("/api")
	// Apply middleware for token verification to the remaining routes
	api.Use(utils.TokenMiddleware)

	api.Post("/createbooks", r.CreateBooks)
	api.Get("/getbooks", r.GetBooks)
	api.Get("/getbook/:id", r.GetBook)
	api.Patch("/updatebook/:id", r.UpdateBook)
	api.Delete("/deletebook/:id", r.DeleteBook)
	// user
}
