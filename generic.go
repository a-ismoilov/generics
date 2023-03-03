package main

type Numbers interface {
	int64 | float64
}

func sumIntsOrFloats[K comparable, V Numbers](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}

	return sum
}
