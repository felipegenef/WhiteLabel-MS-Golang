package main

import (
	"example.com/goLangMicroservice/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/go/api");

	routes.UserRoute(api) //add this

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}