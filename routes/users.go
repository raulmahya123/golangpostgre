// routes/routes.go
package routes

import (
	"golangsidang/repository"
	"golangsidang/utils"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, r *repository.Repositorry) {
	api := app.Group("/api")

	api.Post("/register", r.CreateUser)
	api.Post("/login", r.Login)
	api.Patch("/updateuser/:id", r.UpdateUserByid)
	api.Put("/resetpassword/:username", r.ResetPaswword)

	api.Use(utils.TokenMiddleware)
	api.Get("/getuser", r.GetAllUser)
	api.Get("/getuser/:username", r.GetUserByUsername)
	api.Get("/getuserbyid/:id", r.GetUserById)
	api.Delete("/deleteuser/:id", r.DeleteUser)
}
