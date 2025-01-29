package array_slices

func Sum(numbers []int) (sum int) {
	/* for index := 0; index < 5; index++ {
		sum += numbers[index]
	} */

	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lenOfNumbers := len(numbersToSum)
	sums = make([]int, lenOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
