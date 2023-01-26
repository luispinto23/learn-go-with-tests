package slices

func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
