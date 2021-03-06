package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riririyadi/golang-arch.git/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)

}
