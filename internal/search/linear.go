package search

func LinearSearch(data []int, value int) (bool, int) {

	for i, d := range data {
		if d == value {
			return true, i
		}
	}

	return false, -1
}
