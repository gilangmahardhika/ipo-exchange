package main

import (
	"github.com/gilangmahardhika/ipo-exchange/consumer"
	"github.com/gilangmahardhika/ipo-exchange/model"
)

func main() {
	db := model.DatabaseConnection()

	consumer.ConsumeMessage(db)
}
