package usecase

import (
	"clean-arch/internal/app/repository"
	"clean-arch/internal/domain"
)

type RoomUseCase struct {
	Repo repository.RoomRepository
}

func (uc *RoomUseCase) GetAvailableRooms(status string) ([]domain.Room, error) {
	rooms, err := uc.Repo.FindRoomsByStatus(status)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
