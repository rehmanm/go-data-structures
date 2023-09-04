package trees

import (
	"fmt"
)

type Node struct {
	Data string
}
type Graph struct {
	matrix [][]int8
	nodes  []Node
}

func (g *Graph) New(size int) {
	g.matrix = make([][]int8, size)
	for i := range g.matrix {
		g.matrix[i] = make([]int8, size)
	}
}

func (g *Graph) AddNode(node Node) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) AddEdge(src, dst int, distance int8) {
	g.matrix[src][dst] = distance
}

func (g *Graph) CheckEdge(src, dst int) bool {
	if g.matrix[src][dst] == 1 {
		return true
	} else {
		return false
	}
}

func (g *Graph) Print() {

	fmt.Print("\t")
	for _, node := range g.nodes {
		fmt.Print(node.Data + "\t")
	}
	fmt.Println()

	for i := 0; i < len(g.matrix); i++ {
		fmt.Print(g.nodes[i].Data + "\t")
		for j := 0; j < len(g.matrix[i]); j++ {
			fmt.Print(g.matrix[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Graph) DepthFirstSearch(src int) {
	fmt.Println("Length of Matrix", len(g.matrix))
	visited := make([]bool, len(g.matrix))

	g.dfsHelper(src, visited)
}

func (g *Graph) dfsHelper(src int, visited []bool) {
	if visited[src] {
		return
	} else {
		visited[src] = true
		fmt.Println(g.nodes[src].Data, " = visited")
	}

	for i := 0; i < len(g.matrix[src]); i++ {
		if g.matrix[src][i] > 0 {
			g.dfsHelper(i, visited)
		}
	}

}

func (g *Graph) BreadthFirstSearch(src int) {

	fmt.Println("BreadthFirstSearch")
	visited := make([]bool, len(g.matrix))
	queue := []int{}

	queue = enqueue(queue, src)
	visited[src] = true

	for len(queue) > 0 {
		src, queue = dequeue(queue)
		fmt.Println(g.nodes[src].Data, " = visited")

		for i := 0; i < len(g.matrix[src]); i++ {
			if g.matrix[src][i] > 0 && !visited[i] {
				queue = enqueue(queue, i)
				visited[i] = true
			}
		}
	}
}

func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []int) (int, []int) {
	element := queue[0]
	if len(queue) == 1 {
		tmp := []int{}
		return element, tmp
	}
	return element, queue[1:]
}
