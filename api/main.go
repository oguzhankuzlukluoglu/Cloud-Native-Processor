package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/uploads", "./uploads")

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Dosya seçilmedi"})
		}

		savePath := fmt.Sprintf("./uploads/%s", file.Filename)
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Dosya kaydedilemedi"})
		}

		return c.JSON(fiber.Map{
			"message": "Dosya başarıyla yüklendi",
			"filename": file.Filename,
		})
	})

	log.Fatal(app.Listen(":3000"))
}