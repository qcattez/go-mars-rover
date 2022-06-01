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

func (r *Rover) ExecuteCommands(commands []Command) {
	for i := 0; i < len(commands); i++ {
		r.executeCommand(commands[i])
	}
}

func (r *Rover) executeCommand(command Command) {
	if command == Forward {
		r.moveForward()
	}
	if command == Backward {
		r.moveBackward()
	}
	if command == TurnLeft {
		r.turnLeft()
	}
	if command == TurnRight {
		r.turnRight()
	}
}

func (r *Rover) moveForward() {
	r.position = r.position.next(r.orientation, r.planetCircumference)
}

func (r *Rover) moveBackward() {
	r.position = r.position.next(r.orientation.Opposite(), r.planetCircumference)
}

func (r *Rover) turnLeft() {
	r.orientation = r.orientation.TurnLeft()
}

func (r *Rover) turnRight() {
	r.orientation = r.orientation.TurnRight()
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
