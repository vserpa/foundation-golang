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

	inputUnity := os.Args[len(os.Args)-1]
	inputValues := os.Args[1 : len(os.Args)-1] // slicing ignores the last element

	if inputUnity != "celsius" && inputUnity != "kilometers" {
		fmt.Printf("%s is not a valid unity", inputUnity)
		os.Exit(1)
	}

	for i, v := range inputValues {
		inputValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("Error parsing value '%s' in position %d: %v\n", v, i, err)
			os.Exit(1)
		}

		var outputValue float64
		var outputUnity string
		switch inputUnity {
		case "celsius":
			outputUnity = "fahrenheit"
			outputValue = inputValue*1.8 + 32
		case "kilometers":
			outputUnity = "miles"
			outputValue = inputValue / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", inputValue, inputUnity, outputValue, outputUnity)
	}

}
