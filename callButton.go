package main

// Button on a floor or basement to go back to lobby
type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewCallButton(_id int, _status string, _floor int, _direction string) *CallButton {
	b := new(CallButton)
	b.status = _status
	b.floor = _floor
	b.direction = _direction
	return b
}
