package main

import (
	"os"
)

func main() {
	program := os.Args[1]

	switch program {
	case "converter":
		runConverter()
	case "infinite-loop":
		runInfiniteLoop()
	}
}

func runConverter() {
	converter()
}

func runInfiniteLoop() {
	infiniteLoop()
}
