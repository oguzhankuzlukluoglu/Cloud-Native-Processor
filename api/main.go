package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func initRedis() {
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis at", redisAddr)
}

func main() {
	initRedis()

	app := fiber.New()
	app.Use(logger.New())

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to parse the uploaded file.",
			})
		}

		savePath := fmt.Sprintf("./uploads/%s", file.Filename)
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Could not save the file.",
			})
		}

		err = rdb.LPush(ctx, "task_queue", file.Filename).Err()
		if err != nil {
			log.Printf("Error pushing to Redis: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "File saved but failed to queue for processing.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":   "success",
			"message":  "File uploaded and queued for processing",
			"filename": file.Filename,
		})
	})

	log.Fatal(app.Listen(":3000"))
}