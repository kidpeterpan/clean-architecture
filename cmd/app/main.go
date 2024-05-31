package main

import (
	"context"

	v1 "clean-arch/api/v1"
	"clean-arch/internal/app/repository"
	"clean-arch/internal/app/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://root:toor@localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// use in-memory repository
	//rooms := []domain.Room{
	//	{Code: "1", Size: 200, Price: 100, Longitude: -0.1257, Latitude: 51.5085},
	//}
	//
	//repo := &repository.InMemoryRoomRepository{Rooms: rooms}

	// use MongoDB repository
	repo := &repository.MongoRoomRepository{Client: client}
	useCase := &usecase.RoomUseCase{Repo: repo}
	handler := &v1.Handler{UseCase: useCase}

	r := gin.Default()
	r.GET("/rooms", handler.GetAvailableRooms)
	r.Run(":8080")
}
