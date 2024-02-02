package getOneUser

import (
	global "example.com/goLangMicroservice/Global/Interfaces"
	"github.com/gofiber/fiber/v2"
)

type GetOneUserController struct {
	service *GetOneUserService
}


type ErrorResponse struct {
	Message string `json:"message"`
}

func (controller *GetOneUserController) Handle(ctx *fiber.Ctx) error {
	
	user, err := controller.service.Execute(ctx.Params("id"))
	if err != nil {
		response := ErrorResponse{
			Message: "User not found",
		}
		return ctx.Status(404).JSON(response)
	}



	return ctx.Status(200).JSON(user)
}

func NewGetOneUserController(service *GetOneUserService) global.Controller {
	return &GetOneUserController{
		service,
	}
}
