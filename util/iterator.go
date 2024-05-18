package util

func Iterator(from, to int) (func() int, func() bool) {
	currentVal, step := from, 1
	if from > to {
		step = -1
	}
	done := false

	nextFn := func() int {
		out := currentVal
		if out == to {
			done = true
		} else {
			currentVal += step
		}
		return out
	}

	doneFn := func() bool { return done }

	return nextFn, doneFn
}
