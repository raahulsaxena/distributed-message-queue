package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"distributed-message-queue/internal/rabbitmq"
)

func main() {
	fmt.Println("Starting Publisher Service...")

	// Optional: Validate required env variables
	if os.Getenv("RABBITMQ_URL") == "" || os.Getenv("RABBITMQ_QUEUE") == "" {
		log.Fatal("Missing required environment variables RABBITMQ_URL or RABBITMQ_QUEUE")
	}

	ticker := time.NewTicker(2 * time.Second) // Publish a message every 2 seconds
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			message := fmt.Sprintf("Message published at %s", t.Format(time.RFC3339))
			err := rabbitmq.PublishMessage(message)
			if err != nil {
				log.Printf("Failed to publish message: %v", err)
			}
		}
	}
}