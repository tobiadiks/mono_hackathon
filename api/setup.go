package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var CODE string

type SetupCodeRequestBody struct {
	Code string `json:"code"`
}

func SetupMonoCode(c *fiber.Ctx) error {
	scrb := SetupCodeRequestBody{}

	err := c.BodyParser(&scrb)

	if err != nil {
		c.Status(400).JSON(err.Error())
	}

	CODE = scrb.Code
	fmt.Println(CODE)
	return c.Status(200).SendString("Mono Code has been set")
}

func CheckMonoCode(c *fiber.Ctx) error {
	if CODE == "" {
		return c.Status(400).SendString("Mono code not set")
	}

	return c.Status(200).SendString("Mono code is set")
}

func SetupEndpoints(app *fiber.App) {
	code := app.Group("/code")

	code.Post("/", SetupMonoCode)
	code.Get("/", CheckMonoCode)

}
