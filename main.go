package main

import (
	"fmt"
	"mono-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println(controllers.GetFullAccoutDetails("code_ixj294ZJafLENzdFzjAR"))
	app := fiber.New()
	app.Static("/", "./static")
	// api.SetupEndpoints(app)
	// app.Listen(":3000")
}
