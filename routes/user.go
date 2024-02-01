package routes

import (
	getOneUser "example.com/goLangMicroservice/useCases/User/GetOne"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router) {
	api := app.Group("/users");
	
	api.Get("/:id", getOneUser.Controller.Handle)
}