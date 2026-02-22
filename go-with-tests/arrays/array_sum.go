package arrays

func ArraySum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func SumAll (numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, ArraySum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, ArraySum(tail))
	}

	return sums
}