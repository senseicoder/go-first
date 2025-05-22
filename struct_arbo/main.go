package main

import (
	"fmt"
	"strings"
)

type Node[E any] struct {
	Value E
	Left  *Node[E]
	Right *Node[E]
}

/*
type NodeInt struct {
	Value int
	Left  *NodeInt
	Right *NodeInt
}
*/

var Tronc *Node[int] = &Node[int]{
	Value: 1,
	Left: &Node[int]{
		Value: 2,
		Left: &Node[int]{
			Value: 6,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &Node[int]{
		Value: 3,
		Left: &Node[int]{
			Value: 4,
			Left:  nil,
			Right: nil,
		},
		Right: &Node[int]{
			Value: 5,
			Left:  nil,
			Right: nil,
		},
	},
}

func (n *Node[E]) PrintTree(level int) {
	if n == nil {
		return
	}

	spaces := strings.Repeat("-", level*2)
	fmt.Println(spaces + fmt.Sprint(n.Value))

	n.Left.PrintTree(level + 1)
	n.Right.PrintTree(level + 1)
}

func (n *Node[E]) CountTree() int {
	if n == nil {
		return 0
	}
	return int(1) + n.Left.CountTree() + n.Right.CountTree()
}

func (n *Node[E]) SumTree() int {
	if n == nil {
		return 0
	}
	return any(n.Value).(int) + n.Left.SumTree() + n.Right.SumTree()
}

func main() {
	fmt.Println(Tronc)

	Tronc.PrintTree(0)
	fmt.Println("Total nodes:", Tronc.CountTree())
	fmt.Println("Sum of nodes:", Tronc.SumTree())
}
