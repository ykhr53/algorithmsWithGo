package main

// 210. Course Schedule II
// https://leetcode.com/problems/course-schedule-ii/

var visited map[int]string
var loop bool
var adjList map[int][]int
var ans []int

func findOrder(numCourses int, prerequisites [][]int) []int {
	visited = make(map[int]string)
	adjList = make(map[int][]int)
	ans = make([]int, 0)
	loop = false

	//ans := make([]int, numCourses)

	// create adjList
	for _, pre := range prerequisites {
		adjList[pre[1]] = append(adjList[pre[1]], pre[0])
	}

	// initialise visited map
	for i := 0; i < numCourses; i++ {
		visited[i] = "unvisited"
	}

	// start dfs
	for n := 0; n < numCourses; n++ {
		if visited[n] == "unvisited" {
			dfs(n)
		}
	}
	if loop {
		return make([]int, 0)
	}
	return ans
}

func dfs(n int) {
	if visited[n] == "tmp" {
		loop = true
		return
	} else if visited[n] == "unvisited" {
		visited[n] = "tmp"
		for _, v := range adjList[n] {
			dfs(v)
		}
		visited[n] = "finished"
		ans = append([]int{n}, ans...)
	}
}
