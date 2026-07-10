package main

import (
	"fmt"
	"strings"
)

func stats(words []string) map[string]int {
	stats := make(map[string]int)

	for _, word := range words {
		initial := strings.ToUpper(string(word[0])) // string() converts byte to string
		count, found := stats[initial]
		if found {
			stats[initial] = count + 1
		} else {
			stats[initial] = 1
		}
	}

	return stats
}

func printStats(stats map[string]int) {
	for initial, count := range stats {
		fmt.Printf("Initial: %s, Count: %d\n", initial, count)
	}
}
