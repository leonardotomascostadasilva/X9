package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	KafkaServer string
	HttpPort    string
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
}

func Get() Config {
	return config
}
