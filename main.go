package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Room struct {
	Code      string  `json:"code"`
	Size      int     `json:"size"`
	Price     int     `json:"price"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type RoomRepository interface {
	FindRoomsByStatus(status string) ([]Room, error)
}

type InMemoryRoomRepository struct {
	rooms []Room
}

func (repo *InMemoryRoomRepository) FindRoomsByStatus(status string) ([]Room, error) {
	return repo.rooms, nil
}

type RoomUseCase struct {
	repo RoomRepository
}

func (uc *RoomUseCase) GetAvailableRooms() ([]Room, error) {
	return uc.repo.FindRoomsByStatus("available")
}

type Handler struct {
	useCase *RoomUseCase
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.useCase.GetAvailableRooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rooms)
}

func main() {
	rooms := []Room{
		{Code: "1", Size: 200, Price: 100, Longitude: -0.1257, Latitude: 51.5085},
	}

	repo := &InMemoryRoomRepository{rooms: rooms}
	useCase := &RoomUseCase{repo: repo}
	handler := &Handler{useCase: useCase}

	http.Handle("/rooms", handler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
