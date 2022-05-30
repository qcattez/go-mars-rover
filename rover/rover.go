package rover

import "fmt"

type Rover struct{}

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
		return Rover{}, nil
	default:
		return Rover{}, InvalidOrientationError{orientation}
	}
}

func (r Rover) Move(commands []string) {

}

type Position struct {
	absiss      int
	ordinate    int
	orientation string
}

func (r Rover) Position() Position {
	return Position{0, 0, "N"}
}
