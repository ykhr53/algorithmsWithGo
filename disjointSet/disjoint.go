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
	// run basic union-find
	n := new(Node)
	for i := 1; i < 8; i++ {
		n.makeSet(i)
	}
	n.union(1, 2)
	n.union(2, 3)
	n.union(4, 5)
	n.union(6, 7)
	n.union(3, 7)

	fmt.Println("### main ###")
	for i := 1; i < 8; i++ {
		fmt.Printf("Node %d belongs Node %d\n", i, n.findSet(m[i]).val)
	}

	s := []string{" /", "/ "}
	// run LC 959 problem
	fmt.Println("### LC 959 ###")
	fmt.Printf("lc answer: %d\n", regionsBySlashes(s))
}

func (n *Node) makeSet(val int) {
	node := new(Node)
	node.val = val
	node.rank = 0
	node.parent = node
	m[val] = node
}

func (n *Node) union(val1 int, val2 int) {
	node1 := m[val1]
	node2 := m[val2]

	parent1 := n.findSet(node1)
	parent2 := n.findSet(node2)

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

func (n *Node) findSet(node *Node) *Node {
	parent := node.parent
	if parent == node {
		return parent
	}
	node.parent = n.findSet(node.parent)
	return node.parent
}
