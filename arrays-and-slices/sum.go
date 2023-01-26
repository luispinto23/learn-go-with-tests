package slices

func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	numbersLength := len(numbersToSum)
	sums := make([]int, numbersLength)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
