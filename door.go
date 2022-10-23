package main

type Door struct {
	ID     int
	status string
}

func (requestDoor *Door) NewDoor(_id int, _status string) *Door {
	b := new(Door)
	b.status = _status
	return b
}
