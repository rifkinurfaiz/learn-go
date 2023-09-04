package cube

func Area(numbers ...int) float64 {
	return float64(2 * (numbers[0] * numbers[1] + numbers[1] * numbers[2] + numbers[0] * numbers[2]))
}

type Clojure func(...int) float64

func Volume(height float64, callback Clojure) float64 {
	const width int = 6
	const length int = 7

	return height * callback(5, width, length)
}
