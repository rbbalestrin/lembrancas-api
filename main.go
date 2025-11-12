package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %$", err)
	}
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
