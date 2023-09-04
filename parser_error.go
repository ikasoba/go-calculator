package main

type ParserError struct {
	Msg string
}

func NewParserError(msg string) ParserError {
	return ParserError{msg}
}

func (err ParserError) Error() string {
	return "ParserError: " + err.Msg
}
