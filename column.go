package main

type Column struct {
	ID                int
	status            string
	amountOfFloors    int
	amountOfElevators int
	servedFloorsList  []int
	isBasement        bool
	elevatorsList     []*Elevator
	callButtonList    []CallButton
	bestElevator      *Elevator
	bestScore         int
	referenceGap      int
	direction         string
	currentFloor      int
}

func NewColumn(_id int, _status string, _amountOfFloors int, _amountOfElevators int, _servedFloorsList []int, _isBasement bool) *Column {
	c := new(Column)
	c.status = _status
	c.amountOfFloors = _amountOfFloors
	c.amountOfElevators = _amountOfElevators
	c.callButtonList = make([]CallButton, 0)
	c.servedFloorsList = _servedFloorsList
	c.createElevators(_amountOfFloors, _amountOfElevators)
	c.createCallButtons(_amountOfFloors, _isBasement)
	return c
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	callButtonID := 1
	if _isBasement == true {
		buttonFloor := -1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := NewCallButton(callButtonID, "off", buttonFloor, "up")
			c.callButtonList = append(c.callButtonList, *callButton)
			buttonFloor--
			callButtonID++
		}
	} else {
		for i := 0; i < _amountOfFloors; i++ {
			buttonFloor := -1
			callButton := NewCallButton(callButtonID, "off", buttonFloor, "down")
			c.callButtonList = append(c.callButtonList, *callButton)
			buttonFloor++
			callButtonID++
		}
	}
}

func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	elevatorID := 1
	for i := 0; i < _amountOfElevators; i++ {
		elevator := NewElevator(elevatorID, "online", _amountOfFloors, 1)
		c.elevatorsList = append(c.elevatorsList, elevator)
		elevatorID++
	}
}

// Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator := c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	elevator.addNewRequest(1)
	elevator.move()
	return elevator
}

func (c *Column) findElevator(_requestedFloor int, _requestedDirection string) *Elevator {
	//bestElevator := nil
	bestScore := 6
	referenceGap := 10000000
	if _requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			if elevator.currentFloor == 1 && elevator.status == "stopped" {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			} else if elevator.currentFloor == 1 && elevator.status == "idle" {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			} else if elevator.currentFloor < 1 && elevator.direction == "up" {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			} else if elevator.currentFloor > 1 && elevator.direction == "down" {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			} else if elevator.status == "idle" {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			} else {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			}
		}
	} else {
		for _, elevator := range c.elevatorsList {
			if elevator.currentFloor == _requestedFloor && elevator.status == "stopped" && _requestedDirection == elevator.direction {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			} else if _requestedFloor > elevator.currentFloor && elevator.direction == "up" && _requestedDirection == "up" {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			} else if _requestedFloor < elevator.currentFloor && elevator.direction == "down" && _requestedDirection == "down" {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			} else if elevator.status == "idle" {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			} else {
				c.bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, c.bestElevator, bestScore, referenceGap, _requestedFloor)
			}
		}
	}
	return c.bestElevator
}

func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator *Elevator, bestElevator *Elevator, bestScore int, referenceGap int, floor int) (*Elevator, int, int) {
	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator
		referenceGap = Abs(newElevator.currentFloor - floor)
	} else if bestScore == scoreToCheck {
		gap := Abs(newElevator.currentFloor - floor)
		if referenceGap > gap {
			bestElevator = newElevator
			referenceGap = gap
		}
	}
	return bestElevator, bestScore, referenceGap
}
