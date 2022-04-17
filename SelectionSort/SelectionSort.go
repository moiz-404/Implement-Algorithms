package main

import "fmt"

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d",&d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}

func selectionSort(sliceiItem []int) {
	var n = len(sliceiItem)
	for i := 0; i < n; i++ {
		var minIndex = i
		for j := i+1; j < n; j++ {
			if sliceiItem[j] < sliceiItem[minIndex] {
				minIndex = j
			}
		}
		sliceiItem[i], sliceiItem[minIndex] = sliceiItem[minIndex], sliceiItem[i]		
		
	}
}

func main() {
	
	fmt.Println("Enter input:")
	x := input([]int{}, nil)
	fmt.Println("Output :")
	//      5 8 4 9 7 6 2 1 3
	fmt.Println("unsorted slice items Inputed:", x)
	selectionSort(x)
	fmt.Println("sorted slice items :", x)
}
