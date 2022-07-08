package errors

const (
	MADARA = iota
	NARUTO
)

const (
	NAF = iota
	NVF
)

var WeebErrors = []string{
	MADARA: "We are born to this world as merely children without a clear path or the intellect to know right from wrong.",
	NARUTO: "f you don’t like your destiny, don’t accept it.",
}

var Errors = []string{
	NAF: "not a file to accept",
	NVF: "not a valid flag",
}
