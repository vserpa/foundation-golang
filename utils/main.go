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
	case "named-loop":
		runNamedLoop()
	case "states":
		runStates()
	}
}

func runConverter() {
	converter()
}

func runInfiniteLoop() {
	infiniteLoop()
}

func runNamedLoop() {
	namedLoop()
}

func runStates() {
	states()
}
