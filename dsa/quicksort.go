package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	values := os.Args[1:]
	numbers := make([]int, len(values))

	for i, v := range values {
		number, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%s isn't a valid number \n", v)
			os.Exit(1)
		}
		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}

// Time complexity: average O(n log n), worst case O(n^2)
// Space complexity: average O(n), worst case O(n^2)
func quicksort(numbers []int) []int {

	if len(numbers) <= 1 { // stop condition
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivotValue := n[pivotIndex]

	// remove the pivo (append a slice base with another slice)
	n = append(n[:pivotIndex], n[pivotIndex+1:]...) // ... = include all elements in the base slice

	left, right := partitioning(n, pivotValue)

	return append(
		append(quicksort(left), pivotValue),
		quicksort(right)...,
	)
}

func partitioning(n []int, pivotValue int) ([]int, []int) {
	var left = make([]int, 0)
	var right = make([]int, 0)

	for _, v := range n { // _ = blank identifier
		if v <= pivotValue {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return left, right
}
