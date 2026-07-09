package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Use: <values> <unity: celsius|kilometers>")
		os.Exit(1)
	}

	unity := os.Args[len(os.Args)-1]
	values := os.Args[1 : len(os.Args)-1] // slicing ignores the last element

	if unity != "celsius" && unity != "kilometers" {
		fmt.Printf("%s is not a valid unity", unity)
		os.Exit(1)
	}

	for i, v := range values {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("Error parsing value '%s' in position %d: %v\n", v, i, err)
			os.Exit(1)
		}

		var resultValue float64
		var resultUnity string
		switch unity {
		case "celsius":
			resultUnity = "fahrenheit"
			resultValue = value*1.8 + 32
		case "kilometers":
			resultUnity = "miles"
			resultValue = value / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", value, unity, resultValue, resultUnity)
	}

}
