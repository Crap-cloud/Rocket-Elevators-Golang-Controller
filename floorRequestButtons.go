package main

// FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewFloorRequestButton(_id int, _status string, _floor int, _direction string) *FloorRequestButton {
	b := new(FloorRequestButton)
	b.status = _status
	b.floor = _floor
	b.direction = _direction
	return b
}
