package rover

import (
	. "go-mars-rover/rover/commands"
	. "go-mars-rover/rover/orientations"
)

type Rover struct {
	position            Position
	orientation         Orientation
	planetCircumference uint
}

type Position struct {
	absiss   uint
	ordinate uint
}

func New(absiss uint, ordinate uint, planetCircumference uint, orientation Orientation) Rover {
	return Rover{Position{absiss, ordinate}, orientation, planetCircumference}
}

func (r Rover) Position() Position {
	return r.position
}

func (r Rover) Orientation() Orientation {
	return r.orientation
}

func (r *Rover) ExecuteCommands(commandStrings []CommandString) {
	for i := 0; i < len(commandStrings); i++ {
		executableCommand := ExecutableCommandFactory(commandStrings[i])
		executableCommand.execute(r)
	}
}

func (p Position) next(orientation Orientation, planetCircumference uint) Position {
	if orientation == North {
		return Position{p.absiss, incrementCoordinate(p.ordinate, planetCircumference)}
	}
	if orientation == South {
		return Position{p.absiss, decrementCoordinate(p.ordinate, planetCircumference)}
	}
	if orientation == West {
		return Position{incrementCoordinate(p.absiss, planetCircumference), p.ordinate}
	}
	return Position{decrementCoordinate(p.absiss, planetCircumference), p.ordinate}
}

func incrementCoordinate(coordinate, planetCircumference uint) uint {
	return (coordinate + 1) % planetCircumference
}

func decrementCoordinate(coordinate, planetCircumference uint) uint {
	if coordinate == 0 {
		return planetCircumference - 1
	}
	return coordinate - 1
}

func ExecutableCommandFactory(commandString CommandString) ExecutableCommand {
	if commandString == Forward {
		return ExecutableForward{}
	}
	if commandString == Backward {
		return ExecutableBackward{}
	}
	if commandString == TurnRight {
		return ExecutableTurnRight{}
	}
	return ExecutableTurnLeft{}
}

type ExecutableCommand interface {
	execute(*Rover)
}

type ExecutableForward struct{}
type ExecutableBackward struct{}
type ExecutableTurnRight struct{}
type ExecutableTurnLeft struct{}

func (e ExecutableForward) execute(rover *Rover) {
	rover.position = rover.position.next(rover.orientation, rover.planetCircumference)
}

func (e ExecutableBackward) execute(rover *Rover) {
	rover.position = rover.position.next(rover.orientation.Opposite(), rover.planetCircumference)
}

func (e ExecutableTurnRight) execute(rover *Rover) {
	rover.orientation = rover.orientation.TurnRight()
}

func (e ExecutableTurnLeft) execute(rover *Rover) {
	rover.orientation = rover.orientation.TurnLeft()
}
