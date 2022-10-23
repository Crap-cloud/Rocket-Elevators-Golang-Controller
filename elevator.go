package main

import (
	"sort"
)

type Elevator struct {
	ID                    int
	status                string
	currentFloor          int
	direction             string
	overweight            bool
	amountOfFloors        int
	completedRequestsList []int
	floorRequestsList     []int
	door                  Door
}

func NewElevator(_id int, _status string, _amountOfFloors int, _currentFloor int) *Elevator {
	e := new(Elevator)
	e.status = _status
	e.currentFloor = _currentFloor
	e.amountOfFloors = _amountOfFloors
	return e
}

func (e *Elevator) move() {
	for len(e.floorRequestsList) != 0 {
		destination := e.floorRequestsList[0]
		e.status = "moving"
		if e.direction == "up" {
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.direction == "down" {
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.operateDoors()
		e.floorRequestsList = append(e.floorRequestsList[:0], e.floorRequestsList[(0+1):]...)
		e.completedRequestsList = append(e.completedRequestsList, destination)
	}
	e.status = "idle"
	e.direction = ""
}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Ints(e.floorRequestsList)
	}
	if e.direction == "down" {
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}
}

func (e *Elevator) operateDoors() {
	obstruction := false
	e.door.status = "opened"
	if !e.overweight {
		e.door.status = "closing"
		if !obstruction {
			e.door.status = "closed"
		}
	}
}

func (e *Elevator) addNewRequest(_requestedFloor int) {
	if !contains(e.floorRequestsList, _requestedFloor) {
		e.floorRequestsList = append(e.floorRequestsList, _requestedFloor)
	}
	if e.currentFloor < _requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > _requestedFloor {
		e.direction = "down"
	}
}
