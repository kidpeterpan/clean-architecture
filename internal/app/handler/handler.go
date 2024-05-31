package handler

import (
	"encoding/json"
	"net/http"

	"clean-arch/internal/app/usecase"
)

type Handler struct {
	UseCase *usecase.RoomUseCase
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.UseCase.GetAvailableRooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rooms)
}
