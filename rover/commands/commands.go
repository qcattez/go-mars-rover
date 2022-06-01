package commands

type Command string

var (
	Forward   Command = "f"
	Backward  Command = "b"
	TurnRight Command = "r"
	TurnLeft  Command = "l"
)
