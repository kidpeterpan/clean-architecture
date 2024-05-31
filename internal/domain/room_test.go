package domain_test

import (
	"testing"

	"clean-arch/internal/domain"
)

func TestRoomCreation(t *testing.T) {
	room := domain.Room{
		Code:      "R001",
		Size:      30,
		Price:     1000,
		Longitude: 12.9715987,
		Latitude:  77.5945627,
	}

	if room.Code != "R001" {
		t.Errorf("Room Code was incorrect, got: %s, want: %s.", room.Code, "R001")
	}

	if room.Size != 30 {
		t.Errorf("Room Size was incorrect, got: %d, want: %d.", room.Size, 30)
	}

	if room.Price != 1000 {
		t.Errorf("Room Price was incorrect, got: %d, want: %d.", room.Price, 1000)
	}

	if room.Longitude != 12.9715987 {
		t.Errorf("Room Longitude was incorrect, got: %f, want: %f.", room.Longitude, 12.9715987)
	}

	if room.Latitude != 77.5945627 {
		t.Errorf("Room Latitude was incorrect, got: %f, want: %f.", room.Latitude, 77.5945627)
	}
}
