package main

import (
	"fmt"

	"github.com/rehmanm/go-data-structures/internal/helper"
	"github.com/rehmanm/go-data-structures/internal/search"
	"github.com/rehmanm/go-data-structures/internal/sortdata"
)

func main() {

	valueToFind := 30

	found, position := search.LinearSearch([]int{10, 50, 30, 90}, valueToFind)

	if found {
		fmt.Printf("Found %d at %d\n", valueToFind, position)
	} else {
		fmt.Printf("Not Found %d\n", valueToFind)
	}

	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(search.BinarySearch(items, 1))

	bubbleSortData := helper.GenerateRandomArray(20, 100)
	fmt.Println(bubbleSortData)
	fmt.Println("After Sort")
	sortdata.BubbleSort(bubbleSortData)
	fmt.Println(bubbleSortData)

	fmt.Println("Quick Sort")
	quickSortData := helper.GenerateRandomArray(7, 100)
	fmt.Println(quickSortData)
	fmt.Println("After Sort")
	sortdata.QuickSort(quickSortData, 0, len(quickSortData)-1)
	fmt.Println(quickSortData)

}
