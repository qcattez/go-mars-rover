package rover

import (
	. "go-mars-rover/rover/commands"
	. "go-mars-rover/rover/orientations"
)

type Rover struct {
	absiss              uint
	ordinate            uint
	planetCircumference uint
	orientation         Orientation
}

type Position struct {
	absiss      uint
	ordinate    uint
	orientation Orientation
}

func New(absiss uint, ordinate uint, planetCircumference uint, orientation Orientation) Rover {
	return Rover{absiss, ordinate, planetCircumference, orientation}
}

func (r Rover) Position() Position {
	return Position{r.absiss, r.ordinate, r.orientation}
}

func (r *Rover) ExecuteCommands(commands []Command) {
	for i := 0; i < len(commands); i++ {
		r.executeCommand(commands[i])
	}
}

func (r *Rover) executeCommand(command Command) {
	if r.isCommandToMoveNorth(command) {
		r.moveNorth()
	}
	if r.isCommandToMoveSouth(command) {
		r.moveSouth()
	}
	if r.isCommandToMoveWest(command) {
		r.moveWest()
	}
	if r.isCommandToMoveEast(command) {
		r.moveEast()
	}
	if r.isCommandToTurnEast(command) {
		r.orientation = East
		return
	}
	if r.isCommandToTurnWest(command) {
		r.orientation = West
		return
	}
	if r.isCommandToTurnNorth(command) {
		r.orientation = North
	}
	if r.isCommandToTurnSouth(command) {
		r.orientation = South
	}
}

func (r Rover) isCommandToMoveNorth(command Command) bool {
	return (r.orientation == North && command == Forward) || (r.orientation == South && command == Backward)
}

func (r Rover) isCommandToMoveSouth(command Command) bool {
	return (r.orientation == South && command == Forward) || (r.orientation == North && command == Backward)
}

func (r Rover) isCommandToMoveWest(command Command) bool {
	return (r.orientation == West && command == Forward) || (r.orientation == East && command == Backward)
}

func (r Rover) isCommandToMoveEast(command Command) bool {
	return (r.orientation == East && command == Forward) || (r.orientation == West && command == Backward)
}

func (r *Rover) moveNorth() {
	r.ordinate = moveForward(r.ordinate, r.planetCircumference)
}

func (r *Rover) moveSouth() {
	r.ordinate = moveBackward(r.ordinate, r.planetCircumference)
}

func (r *Rover) moveWest() {
	r.absiss = moveForward(r.absiss, r.planetCircumference)
}

func (r *Rover) moveEast() {
	r.absiss = moveBackward(r.absiss, r.planetCircumference)
}

func moveForward(coordinate, planetCircumference uint) uint {
	return (coordinate + 1) % planetCircumference
}

func moveBackward(coordinate, planetCircumference uint) uint {
	if coordinate == 0 {
		return planetCircumference - 1
	}
	return coordinate - 1
}

func (r Rover) isCommandToTurnEast(command Command) bool {
	return (r.orientation == North && command == TurnRight) || (r.orientation == South && command == TurnLeft)
}

func (r Rover) isCommandToTurnWest(command Command) bool {
	return (r.orientation == North && command == TurnLeft) || (r.orientation == South && command == TurnRight)
}

func (r Rover) isCommandToTurnNorth(command Command) bool {
	return (r.orientation == East && command == TurnLeft) || (r.orientation == West && command == TurnRight)
}

func (r Rover) isCommandToTurnSouth(command Command) bool {
	return (r.orientation == East && command == TurnRight) || (r.orientation == West && command == TurnLeft)
}
