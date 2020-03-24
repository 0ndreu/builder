package myComputer

type MyComputer struct {
	cpu string
	ram int
	hdd int
}

func (pc *MyComputer) CPU() string {
	return pc.cpu
}

func (pc *MyComputer) Ram() int {
	return pc.ram
}

func (pc *MyComputer) HDD() int {
	return pc.hdd
}
