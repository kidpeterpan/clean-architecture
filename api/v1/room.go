package v1

import (
	"net/http"

	"clean-arch/internal/app/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase *usecase.RoomUseCase
}

func (h *Handler) GetAvailableRooms(c *gin.Context) {
	status := c.Query("status")

	rooms, err := h.UseCase.GetAvailableRooms(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToGetAvailableRoomsResponse(rooms))
}
