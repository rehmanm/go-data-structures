package main

import (
	"fmt"
	"os"

	"github.com/rehmanm/go-data-structures/internal/helper"
	"github.com/rehmanm/go-data-structures/internal/search"
	"github.com/rehmanm/go-data-structures/internal/sortdata"
	"github.com/rehmanm/go-data-structures/internal/trees"
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
	sortdata.QuickSort(quickSortData, 0, len(quickSortData)-1)
	fmt.Println(quickSortData)

	fmt.Println("Merge Sort")

	mergeSortData := helper.GenerateRandomArray(10, 100)
	fmt.Println(mergeSortData)
	sortdata.MergeSort(mergeSortData)
	fmt.Println(mergeSortData)

	fmt.Println("Binary Tree")
	tree := &trees.BinaryTree{}
	tree.Insert(10).Insert(8).Insert(12).Insert(7).Insert(11).Insert(13).Insert(9)
	trees.PrintBinaryNode(os.Stdout, tree.GetRoot(), 0, 'M')

	graph := trees.Graph{}
	graph.New(5)

	graph.AddNode(trees.Node{Data: "A"})
	graph.AddNode(trees.Node{Data: "B"})
	graph.AddNode(trees.Node{Data: "C"})
	graph.AddNode(trees.Node{Data: "D"})
	graph.AddNode(trees.Node{Data: "E"})

	graph.AddEdge(0, 1, 1)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 4, 1)
	graph.AddEdge(2, 3, 1)
	graph.AddEdge(2, 4, 1)
	graph.AddEdge(4, 0, 1)
	graph.AddEdge(4, 2, 1)

	graph.Print()
}
