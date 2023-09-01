package sortdata

import "fmt"

func QuickSort(data []int, leftindex, rightindex int) {

	pi := partition(data, leftindex, rightindex)
	if leftindex < pi-1 {
		QuickSort(data, leftindex, pi-1)
	}
	if pi < rightindex {
		QuickSort(data, pi, rightindex)
	}
}

func partition(data []int, leftindex, rightindex int) int {

	pivot := data[leftindex]

	for leftindex <= rightindex {

		for data[leftindex] < pivot {

			leftindex++

		}

		for data[rightindex] > pivot {

			rightindex--

		}

		if leftindex <= rightindex {
			data[leftindex], data[rightindex] = data[rightindex], data[leftindex]
			leftindex++
			rightindex--
		}
		fmt.Println(data, leftindex, rightindex, pivot)
	}

	return leftindex
}

/*
func QuickSort(data []int, leftindex, rightindex int) {
	pivot := data[leftindex]

	low, high := leftindex, rightindex
	fmt.Println("Step1", leftindex, rightindex, low, high)
	for low <= high {
		fmt.Println("Main Loop", low, high, data)
		for data[low] < pivot {
			fmt.Println("low", data[low], pivot, low, data, data[low] < pivot)
			low++
		}

		for data[high] > pivot {
			fmt.Println("high", data[high], pivot, high, data)
			high--
		}

		if low <= high {
			data[low], data[high] = data[high], data[low]
			low++
			high--
		}
	}
	fmt.Println(leftindex, rightindex, low, high)
	if leftindex < high {
		QuickSort(data, leftindex, high)
	}
	if rightindex > low {
		QuickSort(data, low, rightindex)
	}

}


func QuickSort2(data []int) {

	if len(data) < 2 {
		return
	}

	left, right, pivot := 0, len(data)-1, rand.Int()%len(data)

	fmt.Println("step1", data, pivot, data[pivot])

	data[pivot], data[right] = data[right], data[pivot]

	fmt.Println("step2", data, pivot, data[pivot])

	for i, d := range data {
		fmt.Println("step3", i, data[right], left, pivot, data[pivot])
		if d < data[right] {
			data[left], data[i] = data[i], data[left]
			left++
		}

	}

	data[left], data[right] = data[right], data[left]

	fmt.Println("step10", data, left, right)
	QuickSort2(data[:left])
	QuickSort2(data[left+1:])

}

*/
