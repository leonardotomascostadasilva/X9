package repository

import (
	"context"
	"time"

	"github.com/leonardotomascostadasilva/X9/internal/database"
	"github.com/leonardotomascostadasilva/X9/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	UpsertMessage(ctx context.Context, in domain.Message) (domain.Message, error)
	GetMessagesLast30Days(ctx context.Context) ([]domain.Message, error)
}

func UpsertMessage(ctx context.Context, message domain.Message) (domain.Message, error) {

	dbService := database.GetService()
	db := dbService.GetDbInstance().Database("X9")

	_, err := db.Collection("messages").InsertOne(ctx, message)
	if err != nil {
		return domain.Message{}, err
	}

	return message, nil
}

func GetMessagesLast30Days(ctx context.Context) ([]domain.Message, error) {
	dbService := database.GetService()
	db := dbService.GetDbInstance().Database("X9")

	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)

	filter := bson.M{
		"insertedIn": bson.M{
			"$gte": thirtyDaysAgo,
		},
	}

	opts := options.Find().SetSort(bson.D{{Key: "insertedIn", Value: -1}})

	cursor, err := db.Collection("messages").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []domain.Message

	for cursor.Next(ctx) {
		var message domain.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func GetMessagesBySquad(ctx context.Context, squad string) ([]domain.Message, error) {
	dbService := database.GetService()
	db := dbService.GetDbInstance().Database("X9")

	filter := bson.M{
		"squad": squad,
	}

	opts := options.Find().SetSort(bson.D{{Key: "insertedIn", Value: -1}})

	cursor, err := db.Collection("messages").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []domain.Message

	for cursor.Next(ctx) {
		var message domain.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
