package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gilangmahardhika/ipo-exchange/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	model.DatabaseConnection()

	// Connect to RabbitMQ
	conn, err := amqp.Dial(os.Getenv("RABBIT"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Open a channel
	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	msgs, err := channel.Consume(
		"emiten", // queue
		"",       // consumer
		true,     // auto ack
		false,    // exclusive
		false,    // no local
		false,    // no wait
		nil,      //args
	)
	if err != nil {
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
