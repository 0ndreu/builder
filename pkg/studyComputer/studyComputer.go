package studyComputer

import (
	"builder/pkg/builder"
)

//type PCBuilder interface {
//	SetCPU() PCBuilder
//	SetRAM() PCBuilder
//	SetHDD() PCBuilder
//	GetPC() builder.Computer
//}

type studyComputer struct {
	pc builder.Computer
}

func (b *studyComputer) SetCPU() builder.PCBuilder {
	b.pc.CPU("i5")
	return b
}

func (b *studyComputer) SetRAM() builder.PCBuilder {
	b.pc.Ram(4)
	return b
}

func (b *studyComputer) SetHDD() builder.PCBuilder {
	b.pc.HDD(256)
	return b
}

func (b *studyComputer) GetPC() builder.Computer {
	return b.pc
}

func NewStudyComputer(a builder.Computer) builder.PCBuilder {
	return &studyComputer{pc: a}
}
