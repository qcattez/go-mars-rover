package enums

type Movement struct {
	slug string
}

var (
	Unknown  = Movement{}
	Forward  = Movement{"f"}
	Backward = Movement{"b"}
)
