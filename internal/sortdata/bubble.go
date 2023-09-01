package sortdata

func BubbleSort(data []int) {

	var (
		n      = len(data)
		sorted = false
	)

	for !sorted {
		swapped := false

		for i := 0; i < n-1; i++ {
			if data[i] > data[i+1] {
				data[i+1], data[i] = data[i], data[i+1]
				swapped = true
			}
		}

		if !swapped {
			sorted = true
		}
	}
}
