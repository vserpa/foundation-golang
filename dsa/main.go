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
