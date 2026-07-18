package main

import (
	"fmt"
	"math/rand"
	"time"
)

func infiniteLoop() {
	rand.Seed(time.Now().UnixNano()) // using current unix timestamp as seed
	n := 0

	for {
		n++

		i := rand.Intn(4200)
		fmt.Println(i)

		if i%42 == 0 {
			break
		}
	}

	fmt.Printf("Loop terminated after %d iterations\n", n)
}
