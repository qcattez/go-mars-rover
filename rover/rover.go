package rover

import (
	"fmt"
	. "go-mars-rover/rover/enums"
)

type Rover struct {
	absiss      int
	ordinate    int
	orientation string
}

type InvalidOrientationError struct {
	orientation string
}

func (e InvalidOrientationError) Error() string {
	if e.orientation == "" {
		return "The rover orientation is invalid"
	}
	return fmt.Sprintf("\"%s\" is not a valid rover orientation", e.orientation)
}

func New(absiss int, ordinate int, orientation string) (Rover, error) {
	switch orientation {
	case "N", "W", "E", "S":
		return Rover{absiss, ordinate, orientation}, nil
	default:
		return Rover{}, InvalidOrientationError{orientation}
	}
}

func (r *Rover) Move(commands []Movement) {
	for i := 0; i < len(commands); i++ {
		if commands[i] == Forward {
			r.ordinate++
		}
		if commands[i] == Backward {
			r.ordinate--
		}
	}
}

type Position struct {
	absiss      int
	ordinate    int
	orientation string
}

func (r Rover) Position() Position {
	return Position{r.absiss, r.ordinate, r.orientation}
}
