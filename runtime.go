package main

type Runtime struct{}

func (runtime *Runtime) Exec(tree any) (float64, error) {
	switch v := tree.(type) {
	case Operator:
		left, err := runtime.Exec(v.Left)
		if err != nil {
			return 0.0, err
		}

		right, err := runtime.Exec(v.Right)
		if err != nil {
			return 0.0, err
		}

		switch v.Info.Name {
		case Add:
			return left + right, nil
		case Sub:
			return left - right, nil
		case Mul:
			return left * right, nil
		case Div:
			return left / right, nil
		}
	case Number:
		return float64(v), nil
	}

	return 0.0, NewRuntimeError("invalid tree type")
}
