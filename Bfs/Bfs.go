package main

import "fmt"

type queue struct {
	items []int
}

func NewQeue() *queue {
	return &queue{
		items: make([]int, 0),
	}
}

func (q *queue) Add(i int) {
	q.items = append(q.items, i)
}

func (q *queue) Remove() int {
	if len(q.items) == 0 {
		return -1
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item
}

func (q *queue) HasItems() bool {
	return len(q.items) > 0
}

// func NewBFSCommand() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "bfs",
// 		Short: "...",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			runBFS()
// 		},
// 	}

// 	return cmd
// }

// 1 -> 3,2,4
// 2 -> 5
// 3 ->
// 4 -> 7
// 5 -> 6

// DFS
// visits -> 1,4,7,2,5,6,3
// stack ->

// BFS
// visits -> 1,3,2,4,5,7,6
// queue ->

var graph = map[int][]int{
	1: {3, 2, 4},
	2: {5},
	3: {},
	4: {7},
	5: {6},
}

func runBFS() {
	q := NewQeue()
	rootNode := 1
	visited := make([]int, 0)

	q.Add(rootNode)
	for q.HasItems() {
		n := q.Remove()
		visited = append(visited, n)

		adjs := graph[n]
		for _, adj := range adjs {
			q.Add(adj)
		}
	}

	fmt.Println(visited)
}
