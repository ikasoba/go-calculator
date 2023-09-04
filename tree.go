package main

type OperatorPriority int
type OperatorType int

const (
	Add OperatorType = iota
	Sub
	Mul
	Div
)

type OperatorInfo struct {
	Name     OperatorType
	Symbol   string
	Priority OperatorPriority
}

type Operator struct {
	Info  OperatorInfo
	Left  any
	Right any
}

type Number float64
