package main

func sumFloat(m map[string]float64) float64 {
	var sum float64
	for _, i := range m {
		sum += i
	}

	return sum
}

func sumInt(m map[string]int64) int64 {
	var sum int64
	for _, i := range m {
		sum += i
	}

	return sum
}
