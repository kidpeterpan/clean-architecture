package repository

import "clean-arch/internal/domain"

type RoomRepository interface {
	FindRoomsByStatus(status string) ([]domain.Room, error)
}

type InMemoryRoomRepository struct {
	Rooms []domain.Room
}

func (repo *InMemoryRoomRepository) FindRoomsByStatus(status string) ([]domain.Room, error) {
	return repo.Rooms, nil
}
