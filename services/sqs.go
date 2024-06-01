package services

import (
	"log"

	"github.com/igordevopslabs/zapscan-integration/config"
)

func StartSQSConsumer() {
	svc := config.NewSession()

	for {
		messages := config.ReceiveMessages(svc)

		for _, message := range messages {
			log.Printf("Message from queue: %s", *message.Body)
			createScan(*message.Body)
			config.DeleteMessage(svc, message)
		}
	}
}
