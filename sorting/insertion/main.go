package main

import "fmt"

func main() {
	sorted := insertionSort([]int{5, 2, 4, 6, 1, 3})
	fmt.Println(sorted)
}

func insertionSort(numbers []int) []int {
	A := numbers[:]
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i > -1 && A[i] > key {
			A[i+1] = A[i]
			i -= 1
		}
		A[i+1] = key
	}
	return A
}
