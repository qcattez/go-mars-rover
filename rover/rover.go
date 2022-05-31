package rover

import (
	. "go-mars-rover/rover/commands"
	. "go-mars-rover/rover/orientations"
)

type Rover struct {
	absiss      int
	ordinate    int
	orientation Orientation
}

type Position struct {
	absiss      int
	ordinate    int
	orientation Orientation
}

func New(absiss int, ordinate int, orientation Orientation) Rover {
	return Rover{absiss, ordinate, orientation}
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
	r.ordinate++
}

func (r *Rover) moveSouth() {
	r.ordinate--
}

func (r *Rover) moveWest() {
	r.absiss++
}

func (r *Rover) moveEast() {
	r.absiss--
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
