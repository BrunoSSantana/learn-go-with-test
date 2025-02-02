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
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}
