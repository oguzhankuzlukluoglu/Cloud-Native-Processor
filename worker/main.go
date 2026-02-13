package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	fmt.Println("Worker is waiting for tasks...")

	for {
		result, err := rdb.BLPop(ctx, 0, "task_queue").Result()
		if err != nil {
			log.Printf("Error popping task: %v", err)
			continue
		}

		filename := result[1]
		fmt.Printf("Processing task started for file: %s\n", filename)

		time.Sleep(5 * time.Second)

		fmt.Printf("Task COMPLETED for file: %s\n", filename)
	}
}