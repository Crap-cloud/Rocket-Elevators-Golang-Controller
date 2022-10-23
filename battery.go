package main

import (
	"math"
)

type Battery struct {
	ID                        int
	status                    string
	amountOfColumns           int
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	columnsList               []Column
	floorRequestButtonList    []FloorRequestButton
	floor                     int
	servedFloorsList          []int
}

func NewBattery(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) *Battery {
	b := new(Battery)
	b.amountOfColumns = _amountOfColumns
	b.amountOfFloors = _amountOfFloors
	b.amountOfBasements = _amountOfBasements
	b.amountOfElevatorPerColumn = _amountOfElevatorPerColumn

	if _amountOfBasements > 0 {
		b.createBasementRequestButtons(_amountOfFloors)
		b.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	b.createFloorRequestButtons(_amountOfFloors)
	b.createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn)

	return b
}

func (b *Battery) createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {
	columnID := 1
	b.servedFloorsList = make([]int, 0)
	floor := -1
	for i := 0; i < _amountOfBasements; i++ {
		b.servedFloorsList = append(b.servedFloorsList, floor)
		floor--
	}
	column := NewColumn(columnID, "online", _amountOfBasements, _amountOfElevatorPerColumn, b.servedFloorsList, true)
	b.columnsList = append(b.columnsList, *column)
	columnID++
}

func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfElevatorPerColumn int) {
	columnID := 1
	amountOfFloorsPerColumn := int(math.Ceil((float64(_amountOfFloors) / (float64(_amountOfColumns)))))
	b.floor = 1
	for k := 0; k < _amountOfColumns; k++ {
		b.servedFloorsList = make([]int, 0)
		for j := 0; j < amountOfFloorsPerColumn; j++ {
			if b.floor <= _amountOfFloors {
				b.servedFloorsList = append(b.servedFloorsList, b.floor)
				b.floor++
			}
		}
		column := NewColumn(columnID, "online", _amountOfFloors, _amountOfElevatorPerColumn, b.servedFloorsList, false)
		b.columnsList = append(b.columnsList, *column)
		columnID++
	}
}

func (b *Battery) createFloorRequestButtons(_amountOfFloors int) {
	buttonFloor := 1
	floorRequestButtonID := 1
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, "off", buttonFloor, "up")
		b.floorRequestButtonList = append(b.floorRequestButtonList, *floorRequestButton)
		buttonFloor++
		floorRequestButtonID++
	}
}

func (b *Battery) createBasementRequestButtons(_amountOfBasements int) {
	buttonFloor := -1
	floorRequestButtonID := 1
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, "off", buttonFloor, "down")
		b.floorRequestButtonList = append(b.floorRequestButtonList, *floorRequestButton)
		buttonFloor--
		floorRequestButtonID++
	}
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for _, column := range b.columnsList {
		if contains(column.servedFloorsList, _requestedFloor) {
			return &column
		}
	}
	return &b.columnsList[0]
}

// Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return column, elevator
}
