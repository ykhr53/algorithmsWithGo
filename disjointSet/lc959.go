package main

import "fmt"

// 959. Regions Cut By Slashes
// https://leetcode.com/problems/regions-cut-by-slashes/

// DSU is simple struct for union-finding
type DSU struct {
	val    int
	rank   int
	parent *DSU
}

var dm = make(map[int]*DSU)

func (d *DSU) makeSet(val int) {
	dsu := new(DSU)
	dsu.val = val
	dsu.rank = 0
	dsu.parent = dsu
	dm[val] = dsu
}

func (d *DSU) union(val1 int, val2 int) {
	node1 := dm[val1]
	node2 := dm[val2]

	parent1 := d.find(node1)
	parent2 := d.find(node2)

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

func (d *DSU) find(node *DSU) *DSU {
	parent := node.parent
	if parent == node {
		return parent
	}
	node.parent = d.find(node.parent)
	return node.parent
}

func regionsBySlashes(grid []string) int {
	N := len(grid)
	d := new(DSU)
	for i := 0; i < 4*N*N; i++ {
		d.makeSet(i)
	}

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			root := 4 * (r*N + c)
			var val string
			val = string(grid[r][c])
			if val != "\\" {
				d.union(root+0, root+1)
				d.union(root+2, root+3)
			}
			if val != "/" {
				d.union(root+0, root+2)
				d.union(root+1, root+3)
			}

			if r+1 < N {
				// combine bottom and top
				d.union(root+3, (root+4*N)+0)
			}
			if r-1 >= 0 {
				// combine top and bottom
				d.union(root+0, (root-4*N)+3)
			}
			if c+1 < N {
				// combine right and left
				d.union(root+2, (root+4)+1)
			}
			if c-1 >= 0 {
				// combine left and right
				d.union(root+1, (root-4)+2)
			}
		}
	}
	ans := 0
	for i := 0; i < 4*N*N; i++ {
		fmt.Printf("Node %d parent is %d \n", i, d.find(dm[i]).val)
		if d.find(dm[i]).val == i {
			ans++
		}
	}
	return ans
}
