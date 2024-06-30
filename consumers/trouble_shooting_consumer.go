package consumers

import (
	"context"
	"fmt"

	"github.com/leonardotomascostadasilva/X9/internal/config"
	"github.com/segmentio/kafka-go"
)

func TroubleShootingConsumerExecute() {
	config := kafka.ReaderConfig{
		Brokers:  []string{config.Get().KafkaServer},
		Topic:    "trouble-shooting",
		GroupID:  "g1",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(config)

	for {
		message, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Some error occured", err)
			continue
		}
		fmt.Println("Message trouble_shooting_consumer: ", string(message.Value))

	}
}
