package getOneUser

import (
	"github.com/gofiber/fiber/v2"
)

type GetOneUserController struct {
	service GetOneUserService
}

type SuccessResponse struct {
	User UserDTO
}
type ErrorResponse struct {
	Message string
}

func (controller *GetOneUserController) Handle(ctx *fiber.Ctx) error {
	
	user, err := controller.service.Execute(ctx.Params("id"))
	if err != nil {
		response := ErrorResponse{
			Message: "User not found",
		}
		return ctx.Status(404).JSON(response)
	}

	response := SuccessResponse{
		User: user,
	}

	return ctx.Status(200).JSON(response)
}

func NewGetOneUserController(service GetOneUserService) GetOneUserController {
	return GetOneUserController{
		service,
	}
}
