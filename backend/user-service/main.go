package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {


	//Initial fiber instance
	app := fiber.New()

	//create route for fiber
	app.Get("/", func(c*fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	//exception listener
	err:=app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}