package usecase

import (
	"clean-arch/internal/app/repository"
	"clean-arch/internal/domain"
)

type RoomUseCase struct {
	Repo repository.RoomRepository
}

func (uc *RoomUseCase) GetAvailableRooms() ([]domain.Room, error) {
	return uc.Repo.FindRoomsByStatus("available")
}
