package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

type Config struct {
	KafkaServer string
	HttpPort    string
	DbHost      string
	DbPort      string
}

var config Config

func init() {
	configFilePath := "../.env"

	fmt.Println("reading env from: ", configFilePath)

	e := godotenv.Load(configFilePath)
	if e != nil {
		fmt.Println("error loading env: ", e)
		panic(e.Error())
	}
	config.KafkaServer = os.Getenv("KAFKA_SERVER")
	config.HttpPort = os.Getenv("HTTP_PORT")
	config.DbHost = os.Getenv("DB_HOST")
	config.DbPort = os.Getenv("DB_PORT")
}

func Get() Config {
	return config
}

func ConsumerKafka(topic string, groupId string) *kafka.Reader {
	config := kafka.ReaderConfig{
		Brokers:  []string{config.KafkaServer},
		Topic:    topic,
		GroupID:  groupId,
		MaxBytes: 10e6,
	}

	reader := kafka.NewReader(config)

	return reader
}
