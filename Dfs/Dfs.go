package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

//Pop will removea value at th end
//and RETURNs th remove vlue
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func (s *Stack) HasItems() bool {
	return len(s.items) > 0
}

var graph = map[int][]int{
	1: {3, 2, 4},
	2: {5},
	3: {},
	4: {7},
	5: {6},
}

// 1 -> 3,2,4
// 2 -> 5
// 3 ->
// 4 -> 7
// 5 -> 6

// DFS
// visits -> 1,4,7,2,5,6,3
// stack ->

func main() {
	s := NewStack()
	rootNode := 1
	visited := make([]int, 0)

	s.Push(rootNode)
	for s.HasItems() {
		n := s.Pop()
		visited = append(visited, n)

		adjs := graph[n]
		for _, adj := range adjs {
			s.Push(adj)
		}
	}

	fmt.Println(visited)
}
