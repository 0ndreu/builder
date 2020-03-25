package main

import (
	"builder/pkg/builder"
	"builder/pkg/computer"
	"builder/pkg/gamingComputer"
	//"builder/pkg/builder"
	//"builder/pkg/studyComputer"
	"fmt"
)

func main() {
	machine := computer.NewComputer()
	b := gamingComputer.NewGamingComputer(machine)
	director := builder.NewDirector(b)
	director.ConstructPC()
	fmt.Println(b)
}
