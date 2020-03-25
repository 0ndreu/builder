package gamingComputer

import (
	"builder/pkg/builder"
)

//type PCBuilder interface {
//	SetCPU() PCBuilder
//	SetRAM() PCBuilder
//	SetHDD() PCBuilder
//	GetPC() builder.Computer
//}

type gamingComputer struct {
	pc builder.Computer
}

func (b *gamingComputer) SetCPU() builder.PCBuilder {
	b.pc.CPU("i7")
	return b
}

func (b *gamingComputer) SetRAM() builder.PCBuilder {
	b.pc.Ram(32)
	return b
}

func (b *gamingComputer) SetHDD() builder.PCBuilder {
	b.pc.HDD(1024)
	return b
}

func (b *gamingComputer) GetPC() builder.Computer {
	return b.pc
}

func NewGamingComputer(a builder.Computer) builder.PCBuilder {
	return &gamingComputer{pc: a}
}
