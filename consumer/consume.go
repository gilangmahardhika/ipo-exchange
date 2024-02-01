package consumer

import (
	"encoding/json"
	"log"

	"github.com/gilangmahardhika/ipo-exchange/helper"
	"github.com/gilangmahardhika/ipo-exchange/model"
	"github.com/gilangmahardhika/ipo-exchange/repository"
	"gorm.io/gorm"
)

func ConsumeMessage(db *gorm.DB) {
	conn, _ := rabbitConnection()
	channel, err := conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")
	// defer channel.Close()
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
	defer conn.Close()
	repo := repository.NewIpoRepository(db)

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Println("Receiving data")
			data := ParsingData(string(msg.Body))
			repo.CreateOrUpdateIPO(data)
		}
	}()

	log.Println("Waiting for messages...")
	<-forever
}

func ParsingData(data string) (Ipo *model.Ipo) {
	json.Unmarshal([]byte(data), &Ipo)
	return Ipo
}
