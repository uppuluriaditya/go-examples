package misc

func FindSecondLargestNumber(array []int) int {
	max, secondMax := -1, -1

	for idx := range array {
		if array[idx] > max {
			secondMax = max
			max = array[idx]
		} else if array[idx] > secondMax && array[idx] != max {
			secondMax = array[idx]
		}
	}
	return secondMax
}
