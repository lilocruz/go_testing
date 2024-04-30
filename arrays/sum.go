package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, nunumber := range numbers {
		sum += nunumber
	}
	return sum
}
