package consumers

import (
	"context"
	"fmt"

	"github.com/leonardotomascostadasilva/X9/internal/config"
)

func TroubleShootingConsumerExecute() {

	reader := config.ConsumerKafka("trouble-shooting", "g1")

	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Some error occured", err)
			continue
		}
		fmt.Println("Message trouble_shooting_consumer: ", string(message.Value))

	}
}
