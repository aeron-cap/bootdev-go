package arrays

import "math"

func sortEven(numbers []int) []int {
	var even []int
	for _, number := range numbers {
		if (math.Mod(float64(number), 2) == 0) {
			even = append(even, number)
		}
	}

	return even
}

func sortOdd(numbers []int) []int {
	var odd []int
	for _, number := range numbers {
		if (math.Mod(float64(number), 2) != 0) {
			odd = append(odd, number)
		}
	}

	return odd
}

func sortArray(numberlist ...[]int) [][]int {
	var numberArray [][]int

	for _, numbers := range numberlist {
		numberArray = append(numberArray, sortEven(numbers), sortOdd(numbers))
	}

	return numberArray
}