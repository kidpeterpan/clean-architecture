package repository

import (
	"context"
	"fmt"

	"clean-arch/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository interface {
	FindRoomsByStatus(status string) ([]domain.Room, error)
}

type InMemoryRoomRepository struct {
	Rooms []domain.Room
}

func (repo *InMemoryRoomRepository) FindRoomsByStatus(status string) ([]domain.Room, error) {
	if status != "available" {
		return nil, fmt.Errorf("invalid status")
	}
	return repo.Rooms, nil
}

type MongoRoomRepository struct {
	Client *mongo.Client
}

func (repo *MongoRoomRepository) FindRoomsByStatus(status string) ([]domain.Room, error) {
	if status != "available" {
		return nil, fmt.Errorf("invalid status")
	}

	cur, err := repo.Client.Database("rooms").Collection("rooms").Find(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var rooms []domain.Room
	for cur.Next(context.TODO()) {
		var room domain.Room
		err := cur.Decode(&room)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
