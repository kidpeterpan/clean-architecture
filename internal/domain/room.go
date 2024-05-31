package domain

type Room struct {
	Code      string  `bson:"code"`
	Size      int     `bson:"size"`
	Price     int     `bson:"price"`
	Longitude float64 `bson:"longitude"`
	Latitude  float64 `bson:"latitude"`
}
