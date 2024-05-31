package main

import (
	"fmt"
	"net/http"

	"clean-arch/internal/app/handler"
	"clean-arch/internal/app/repository"
	"clean-arch/internal/app/usecase"
	"clean-arch/internal/domain"
)

func main() {
	rooms := []domain.Room{
		{Code: "1", Size: 200, Price: 100, Longitude: -0.1257, Latitude: 51.5085},
	}

	repo := &repository.InMemoryRoomRepository{Rooms: rooms}
	useCase := &usecase.RoomUseCase{Repo: repo}
	h := &handler.Handler{UseCase: useCase}

	http.Handle("/rooms", h)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
