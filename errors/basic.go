package errors

const (
	NAF = iota
	NVF
)

var Errors = []string{
	NAF: "not a file to accept",
	NVF: "not a valid flag",
}
