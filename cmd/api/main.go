package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/raven870/internal/config"
)

func main() {
	fmt.Println("This is the Test application running the api service !!!")

	//1. take the variable from the Enivronment
	cfg := config.ConfigureEnv()
	Port := cfg.PORT
	mode := cfg.Env
	//2. Configure the Logger for the Structured Logging in the production
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	//3. Now start the Fiber Server

	app := fiber.New(fiber.Config{
		DisableDefaultContentType: true,
		AppName:                   strings.ToUpper(mode),
	})
	//make a health test
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "api service is running",
			"mode":    mode,
		})
	})

	app.Get("/user", func(c fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data": fiber.Map{
				"username": "gautam",
				"password": "123456",
				"address":  "jnk-04",
			},
		})
	})

	err := app.Listen(":" + Port)
	if err != nil {
		slog.Debug("Error while starting server ->", "ErrorMsz", err.Error())
	}

}
