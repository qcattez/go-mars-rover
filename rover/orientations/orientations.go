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
