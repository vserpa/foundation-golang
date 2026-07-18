package main

import "fmt"

func namedLoop() {
	var i int

loop:
	for i = 0; i < 10; i++ {
		fmt.Printf("for i = %d\n", i)

		switch i {
		case 2, 3:
			fmt.Printf("breaking switch, i == %d\n", i)
			break
		case 5:
			fmt.Printf("breaking loop, i == %d\n", i)
			break loop
		}
	}
}
