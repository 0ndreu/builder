package main

import (
	"builder/pkg/builder"
	"builder/pkg/computer"
	"builder/pkg/gamingComputer"
	//"builder/pkg/studyComputer"
	"fmt"
)

func main() {
	gamingMachine := computer.NewComputer()
	a := gamingComputer.NewGamingComputer(gamingMachine)
	directorGame := builder.NewDirector(a)
	directorGame.ConstructPC()
	ready := a.GetPC()
	fmt.Println(ready)
}
