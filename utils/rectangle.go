package utils

func Area(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return float64(total)
}
