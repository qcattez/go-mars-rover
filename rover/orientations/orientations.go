package orientations

type Orientation struct {
	slug string
}

var (
	Unknown = Orientation{}
	North   = Orientation{"N"}
	East    = Orientation{"E"}
	West    = Orientation{"W"}
	South   = Orientation{"S"}
)

type InvalidOrientationError struct{}

func (e InvalidOrientationError) Error() string {
	return "The rover orientation is unknown"
}
