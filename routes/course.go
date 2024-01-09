// routes/routes.go
package routes

import (
	"golangsidang/repository"
	"golangsidang/utils"

	"github.com/gofiber/fiber/v2"
)

func CourseRoutes(app *fiber.App, r *repository.CourseRepository) {
	api := app.Group("/api")
	api.Use(utils.TokenMiddleware)
	// Apply middleware for token verification to the remaining routes
	api.Post("/createcourse", r.CreateCourse)
	api.Get("/getcourse", r.GetCourses)
	api.Get("/getcourse/:id", r.GetCourseByid)
	api.Patch("/updatecourse/:id", r.UpdateCourse)
	api.Delete("/deletecourse/:id", r.DeleteCourse)
}
