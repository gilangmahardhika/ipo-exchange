package consumer

import (
	"os"

	"github.com/gilangmahardhika/ipo-exchange/helper"
	amqp "github.com/rabbitmq/amqp091-go"
)

func rabbitConnection() (conn *amqp.Connection, err error) {
	conn, error := amqp.Dial(os.Getenv("RABBIT"))
	helper.FailOnError(error, "Failed to connect to RabbitMQ")
	return conn, err
}
