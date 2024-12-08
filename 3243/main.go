package main

import "fmt"

func main() {
	n := 5
	queries := [][]int{{2, 4}, {0, 2}, {0, 4}}
	fmt.Println(shortestDistanceAfterQueries(n, queries)) // [3, 2, 1]

	fmt.Println("----------------------------")
	n = 4
	queries = [][]int{{0, 3}, {0, 2}}
	fmt.Println(shortestDistanceAfterQueries(n, queries)) // [1, 1]

	fmt.Println("----------------------------")
	n = 5
	queries = [][]int{{1, 4}, {1, 3}}
	fmt.Println(shortestDistanceAfterQueries(n, queries)) // [2, 2]

	fmt.Println("----------------------------")
	n = 6
	queries = [][]int{{1, 4}, {0, 2}}
	fmt.Println(shortestDistanceAfterQueries(n, queries)) // [3, 3]

	fmt.Println("----------------------------")
	n = 7
	queries = [][]int{{4, 6}, {0, 3}}
	fmt.Println(shortestDistanceAfterQueries(n, queries)) // [5, 3]
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	res := make([]int, len(queries))

	conections := make(map[int][]int)

	for i := 0; i < n-1; i++ {
		conections[i] = []int{i + 1}
	}
	for i := 0; i < len(queries); i++ {
		conections[queries[i][0]] = append(conections[queries[i][0]], queries[i][1])
		res[i] = bfsDistance(conections, 0, n-1)
	}

	return res
}

func bfsDistance(graph map[int][]int, start, end int) int {
	queue := []int{start}
	visited := make(map[int]bool)
	distance := make(map[int]int)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == end {
			break
		}

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				distance[neighbor] = distance[node] + 1
			}
		}
	}

	return distance[end]
}
