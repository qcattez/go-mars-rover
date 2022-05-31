package commands

type Command struct {
	slug string
}

var (
	Unknown   = Command{}
	Forward   = Command{"f"}
	Backward  = Command{"b"}
	TurnRight = Command{"r"}
	TurnLeft  = Command{"l"}
)
