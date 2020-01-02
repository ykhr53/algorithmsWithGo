package main

import "fmt"

// Node data structure
type Node struct {
	val    int
	rank   int
	parent *Node
}

var m = make(map[int]*Node)

func main() {
	for i := 1; i < 8; i++ {
		makeSet(i)
	}
	union(1, 2)
	union(2, 3)
	union(4, 5)
	union(6, 7)
	union(3, 7)

	for i := 1; i < 8; i++ {
		fmt.Printf("Node %d belongs Node %d\n", i, findSet(m[i]).val)
	}
}

func makeSet(val int) {
	node := new(Node)
	node.val = val
	node.rank = 0
	node.parent = node
	m[val] = node
}

func union(val1 int, val2 int) {
	node1 := m[val1]
	node2 := m[val2]

	parent1 := findSet(node1)
	parent2 := findSet(node2)

	if parent1.val == parent2.val {
		return
	}

	if parent1.rank >= parent2.rank {
		if parent1.rank == parent2.rank {
			parent1.rank++
		}
		parent2.parent = parent1
	} else {
		parent1.parent = parent2
	}
}

func findSet(node *Node) *Node {
	parent := node.parent
	if parent == node {
		return parent
	}
	node.parent = findSet(node.parent)
	return node.parent
}
