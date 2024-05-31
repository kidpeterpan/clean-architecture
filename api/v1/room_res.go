package v1

import "clean-arch/internal/domain"

type GetAvailableRoomsResponse struct {
	Rooms []RoomResponse `json:"rooms"`
}

type RoomResponse struct {
	Code      string  `json:"code"`
	Size      int     `json:"size"`
	Price     int     `json:"price"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func ToGetAvailableRoomsResponse(room []domain.Room) GetAvailableRoomsResponse {
	var res GetAvailableRoomsResponse
	for _, r := range room {
		res.Rooms = append(res.Rooms, RoomResponse{
			Code:      r.Code,
			Size:      r.Size,
			Price:     r.Price,
			Longitude: r.Longitude,
			Latitude:  r.Latitude,
		})
	}
	return res
}
