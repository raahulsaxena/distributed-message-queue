package main

import (
	"fmt"
	"log"
	"os"

	"distributed-message-queue/internal/rabbitmq"
)

func main() {
	fmt.Println("Starting Consumer Service...")

	// Optional: Validate required env variables
	if os.Getenv("RABBITMQ_URL") == "" || os.Getenv("RABBITMQ_QUEUE") == "" {
		log.Fatal("Missing required environment variables RABBITMQ_URL or RABBITMQ_QUEUE")
	}

	err := rabbitmq.ConsumeMessages()
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}
}