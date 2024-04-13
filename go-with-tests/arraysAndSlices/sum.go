package main

func main() {

}

//	func SumAllTails(numbersToSum ...[]int) []int {
//		var sums []int
//		for _, numbers := range numbersToSum {
//			tail := numbers[1:]
//			sums = append(sums, Sum(tail))
//		}
//		return sums
//	}
func SumAllTails(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		if len(numbers) < 2 {
			sums[i] = 0
		} else {
			sums[i] = Sum(numbers[1:])
		}
	}
	return sums
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	var sums []int

	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
