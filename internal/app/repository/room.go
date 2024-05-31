package repository

import (
	"fmt"

	"clean-arch/internal/domain"
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
