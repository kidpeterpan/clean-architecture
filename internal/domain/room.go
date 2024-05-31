package domain

type Room struct {
	Code      string  `json:"code"`
	Size      int     `json:"size"`
	Price     int     `json:"price"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
