package trees

import "fmt"

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
