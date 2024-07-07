package database

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/leonardotomascostadasilva/X9/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once     sync.Once
	instance Service
)

type Service interface {
	Health() map[string]string
	GetDbInstance() *mongo.Client
}

type service struct {
	db *mongo.Client
}

func GetService() Service {
	once.Do(func() {
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", config.Get().DbHost, config.Get().DbPort)))
		if err != nil {
			log.Fatal(err)
		}
		instance = &service{
			db: client,
		}
	})
	return instance
}

func (s *service) GetDbInstance() *mongo.Client {
	return s.db
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "Mongo Db it's healthy",
	}
}
