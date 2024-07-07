package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/leonardotomascostadasilva/X9/internal/config"
	"github.com/leonardotomascostadasilva/X9/internal/domain"
	repository "github.com/leonardotomascostadasilva/X9/internal/repositories"
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

		var msg domain.Message

		err = json.Unmarshal(message.Value, &msg)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
			continue
		}

		msg.InsertedIn = time.Now()

		ctx := context.Background()

		createdMessage, err := repository.UpsertMessage(ctx, msg)

		if err != nil {
			log.Fatalf("Erro ao criar mensagem: %v", err)
		}

		fmt.Println("Message trouble_shooting_consumer: ", createdMessage)

	}
}
