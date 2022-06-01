package orientations

type Orientation struct {
	slug string
}

var (
	North = Orientation{"N"}
	East  = Orientation{"E"}
	West  = Orientation{"W"}
	South = Orientation{"S"}
)

func (o Orientation) Opposite() Orientation {
	if o == North {
		return South
	}
	if o == South {
		return North
	}
	if o == West {
		return East
	}
	return West
}

func (o Orientation) TurnLeft() Orientation {
	if o == North {
		return West
	}
	if o == South {
		return East
	}
	if o == West {
		return South
	}
	return North
}

func (o Orientation) TurnRight() Orientation {
	return o.TurnLeft().Opposite()
}
