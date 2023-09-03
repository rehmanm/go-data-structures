package trees

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	leftNode  *BinaryNode
	rightNode *BinaryNode
	data      int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) GetRoot() *BinaryNode {
	return t.root
}

func (t *BinaryTree) Insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, leftNode: nil, rightNode: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.leftNode == nil {
			n.leftNode = &BinaryNode{data: data, leftNode: nil, rightNode: nil}
		} else {
			n.leftNode.insert(data)
		}
	} else {
		if n.rightNode == nil {
			n.rightNode = &BinaryNode{data: data, leftNode: nil, rightNode: nil}
		} else {
			n.rightNode.insert(data)
		}
	}

}

func PrintBinaryNode(w io.Writer, node *BinaryNode, ns int, ch rune) {

	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	PrintBinaryNode(w, node.leftNode, ns+2, 'L')
	PrintBinaryNode(w, node.rightNode, ns+2, 'R')
}
