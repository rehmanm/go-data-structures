package helper

import (
	"math/rand"
)

func GenerateRandomArray(size, max int) []int {

	array := make([]int, size)
	for i := 0; i < size; i++ {

		array[i] = rand.Intn(max)

	}

	return array
}
