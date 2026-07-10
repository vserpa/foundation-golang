package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	program := os.Args[1]

	switch program {
	case "quicksort":
		runQuicksort()
	case "maps":
		runMaps()
	case "stack":
		runStack()
	}

}

func runQuicksort() {
	values := os.Args[2:]
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

func runMaps() {
	words := os.Args[2:]
	stats := stats(words)
	printStats(stats)
}

func runStack() {
	stack := Stack{}
	fmt.Println("Stack created with size:", stack.Size())
	fmt.Println("Is stack empty?", stack.IsEmpty())

	stack.Push("Go")
	stack.Push(2009)
	stack.Push(3.14)
	stack.Push("End")

	fmt.Println("Stack size:", stack.Size())
	fmt.Println("Is stack empty? ", stack.IsEmpty())

	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		fmt.Println("Popped value:", v)
		fmt.Println("Size after pop:", stack.Size())
		fmt.Println("Is stack empty? ", stack.IsEmpty())
	}

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error occurred while popping from stack:", err)
	}
}
