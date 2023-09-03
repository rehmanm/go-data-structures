package sortdata

func MergeSort(data []int) {

	if len(data) <= 1 {
		return
	}

	middle := len(data) / 2

	leftData := make([]int, middle)
	rightData := make([]int, len(data)-middle)

	for i := 0; i < len(data); i++ {
		if i < middle {
			leftData[i] = data[i]
		} else {
			rightData[i-middle] = data[i]
		}
	}

	MergeSort(leftData)
	MergeSort(rightData)

	merge(leftData, rightData, data)

}

func merge(leftData, rightData, data []int) {

	leftSize := len(data) / 2
	rightSize := len(data) - leftSize
	l, r, i := 0, 0, 0
	for l < leftSize && r < rightSize {
		if leftData[l] < rightData[r] {
			data[i] = leftData[l]
			l++
			i++
		} else {
			data[i] = rightData[r]
			r++
			i++
		}
	}

	for l < leftSize {
		data[i] = leftData[l]
		i++
		l++
	}
	for r < rightSize {

		data[i] = rightData[r]
		i++
		r++
	}

}
