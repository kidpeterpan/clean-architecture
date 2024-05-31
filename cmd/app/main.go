package main

import (
	v1 "clean-arch/api/v1"
	"clean-arch/internal/app/repository"
	"clean-arch/internal/app/usecase"
	"clean-arch/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	rooms := []domain.Room{
		{Code: "1", Size: 200, Price: 100, Longitude: -0.1257, Latitude: 51.5085},
	}

	repo := &repository.InMemoryRoomRepository{Rooms: rooms}
	useCase := &usecase.RoomUseCase{Repo: repo}
	handler := &v1.Handler{UseCase: useCase}

	r := gin.Default()
	r.GET("/rooms", handler.GetAvailableRooms)
	r.Run(":8080")
}
