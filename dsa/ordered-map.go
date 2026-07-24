package main

import (
	"fmt"
	"sort"
)

func orderedMap() {
	unorderedMap := make(map[int]int, 15)

	for i := 1; i <= 15; i++ {
		unorderedMap[i] = i * i
	}

	supportSlice := make([]int, 0, len(unorderedMap))

	for key := range unorderedMap {
		supportSlice = append(supportSlice, key)
	}

	sort.Ints(supportSlice)

	for _, key := range supportSlice {
		value := unorderedMap[key]
		fmt.Printf("%d^2 = %d\n", key, value)
	}
}
