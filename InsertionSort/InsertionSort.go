package main

import "fmt"

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}

func insertionSort(sliceiItem []int) {
	var n = len(sliceiItem)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if sliceiItem[j-1] > sliceiItem[j] {
				sliceiItem[j-1], sliceiItem[j] = sliceiItem[j], sliceiItem[j-1]
			}
			j = j - 1
		}
	}
}

func main() {

	fmt.Println("Enter input:")
	x := input([]int{}, nil)
	fmt.Println("Output :")
	fmt.Println("unsorted slice items Inputed:", x)
	insertionSort(x)
	fmt.Println("sorted slice items :", x)
}
