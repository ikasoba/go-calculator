package main

import (
	"regexp"
	"strconv"
)

func ParseNumber(i int, src string) (Number, int, error) {
	if i >= len(src) {
		return Number(0.0), i, NewParserError("index out of range.")
	}

	var tmp string

	for ; i < len(src) && src[i] >= '0' && src[i] <= '9'; i++ {
		tmp += string(src[i])
	}

	if i < len(src) && src[i] != '.' {
		f, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			return Number(0.0), i, err
		}

		return Number(f), i, nil
	}

	for ; i < len(src) && src[i] >= '0' && src[i] <= '9'; i++ {
		tmp += string(src[i])
	}

	f, err := strconv.ParseFloat(tmp, 64)
	if err != nil {
		return Number(0.0), i, err
	}

	return Number(f), i, nil
}

func MatchWhitespace(i int, src string) int {
	for i < len(src) && regexp.MustCompile(`\s`).MatchString(string(src[i])) {
		i++
	}

	return i
}

func ParseOperand(i int, src string) (any, int, error) {
	f, i, err := ParseNumber(i, src)
	if err == nil {
		return f, i, nil
	}

	return nil, i, err
}

func ParseOperator(info OperatorInfo, i int, src string) (*Operator, int, error) {
	if i >= len(src) {
		return nil, i, NewParserError("index out of range.")
	}

	left, i, err := ParseOperand(i, src)
	if err != nil {
		return nil, i, err
	}

	i = MatchWhitespace(i, src)

	if !(i+len(info.Symbol) < len(src) && src[i:i+len(info.Symbol)] == info.Symbol) {
		return nil, i, NewParserError("cannot match \"" + info.Symbol + "\"")
	}

	i += len(info.Symbol)

	i = MatchWhitespace(i, src)

	right, i, err := ParseExpr(i, src)
	if err != nil {
		return nil, i, err
	}

	if v, ok := right.(Operator); ok {
		if info.Priority > v.Info.Priority {
			return &Operator{
				v.Info,
				Operator{
					info,
					left,
					v.Left,
				},
				v.Right,
			}, i, nil
		}
	}

	return &Operator{
		info,
		left,
		right,
	}, i, nil
}

var operators = [...]OperatorInfo{
	{
		Add,
		"+",
		OperatorPriority(0),
	},
	{
		Sub,
		"-",
		OperatorPriority(0),
	},
	{
		Mul,
		"*",
		OperatorPriority(1),
	},
	{
		Div,
		"/",
		OperatorPriority(1),
	},
}

func ParseExpr(i int, src string) (any, int, error) {
	if i >= len(src) {
		return nil, i, NewParserError("index out of range.")
	}

	for _, op := range operators {
		tree, index, err := ParseOperator(op, i, src)
		if err == nil {
			return *tree, index, nil
		}
	}

	return ParseOperand(i, src)
}
