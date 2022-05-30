package rover

import (
	mv "go-mars-rover/rover/movements"
	or "go-mars-rover/rover/orientations"
)

type Rover struct {
	absiss      int
	ordinate    int
	orientation or.Orientation
}

func New(absiss int, ordinate int, orientation or.Orientation) (Rover, error) {
	switch orientation {
	case or.North, or.East, or.West, or.South:
		return Rover{absiss, ordinate, orientation}, nil
	default:
		return Rover{}, or.InvalidOrientationError{}
	}
}

func (r *Rover) Move(commands []mv.Movement) {
	for i := 0; i < len(commands); i++ {
		if commands[i] == mv.Forward {
			r.ordinate++
		}
		if commands[i] == mv.Backward {
			r.ordinate--
		}
	}
}

type Position struct {
	absiss      int
	ordinate    int
	orientation or.Orientation
}

func (r Rover) Position() Position {
	return Position{r.absiss, r.ordinate, r.orientation}
}
