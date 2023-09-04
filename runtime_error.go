package main

type RuntimeError struct {
	Msg string
}

func NewRuntimeError(msg string) RuntimeError {
	return RuntimeError{msg}
}

func (err RuntimeError) Error() string {
	return "RuntimeError: " + err.Msg
}
