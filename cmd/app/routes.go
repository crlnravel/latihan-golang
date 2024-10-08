package main

import (
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type config struct {
	db *gorm.DB
}

func NewApp(cfg *config) *fiber.App {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		if err := c.JSON(fiber.Map{
			"message": "🐣 v1",
		}); err != nil {
			return err
		}
		return c.Next()
	})

	// Tryout routes examples
	//exam.App(v1, cfg.db)
	//subject.App(v1, cfg.db)
	//question.App(v1, cfg.db)
	//option.App(v1, cfg.db)

	return app
}
