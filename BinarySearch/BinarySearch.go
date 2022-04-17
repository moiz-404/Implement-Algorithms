package main

import "fmt"

func binarySearch(arr []int, key int) int {
    leftPointer := 0 //array starting point
    var rightPointer = len(arr) - 1 //array length
    for leftPointer <= rightPointer {
        var midPointer = int((leftPointer + rightPointer) / 2)
        var midValue = arr[midPointer]

        if midValue == key {
            return midPointer
        } else if midValue < key {
            leftPointer = midPointer + 1
        } else {
            rightPointer = midPointer - 1
        }
    }
    return -1
}

func binarySearchs(arr []int, key int) bool {
	low := 0 //array starting point
	high := len(arr) - 1 //array length

	for low <= high{
		midPoint :=int((low + high) / 2)
		if arr[midPoint] < key {
			low = midPoint + 1
		}else{
			high = midPoint - 1
		}
	}

	if low == len(arr) || arr[low] != key {
		return false
	}

	return true
}

func main() {
    var n = []int{2, 9, 11, 21, 22, 32, 36, 48, 76}
    fmt.Println("binary Searched element index : ",binarySearch(n, 32))
	fmt.Println("binary Searched element : ",binarySearchs(n, 32))
}