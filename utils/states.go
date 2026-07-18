package main

import "fmt"

type State struct {
	name       string
	population int
	capital    string
}

func states() {
	states := make(map[string]State, 6)
	states["CA"] = State{name: "California", population: 39538223, capital: "Sacramento"}
	states["TX"] = State{name: "Texas", population: 29145505, capital: "Austin"}
	states["FL"] = State{name: "Florida", population: 21538187, capital: "Tallahassee"}
	states["NY"] = State{name: "New York", population: 20201249, capital: "Albany"}
	states["PA"] = State{name: "Pennsylvania", population: 13002700, capital: "Harrisburg"}
	states["IL"] = State{name: "Illinois", population: 12812508, capital: "Springfield"}

	fmt.Println(states)
}
