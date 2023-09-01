package search

func BinarySearch(data []int, value int) (bool, int) {

	low := 0
	high := len(data) - 1

	for low <= high {

		median := (low + high) / 2

		if data[median] < value {
			low = median + 1
		} else {
			high = median - 1
		}

	}

	if low == len(data) || data[low] != value {
		return false, -1
	}

	return true, low
}
