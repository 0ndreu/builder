package gamingComputer

type computer interface {
	CPU(proc string)
	Ram(size int)
	HDD(size int)
}

type PCBuilder interface {
	SetCPU() PCBuilder
	SetRAM() PCBuilder
	SetHDD() PCBuilder
	GetPC() computer
}

type gamingComputer struct {
	pc computer
}

func (b *gamingComputer) SetCPU() PCBuilder {
	b.pc.CPU("i3")
	return b
}

func (b *gamingComputer) SetRAM() PCBuilder {
	b.pc.Ram(8)
	return b
}

func (b *gamingComputer) SetHDD() PCBuilder {
	b.pc.HDD(512)
	return b
}

func (b *gamingComputer) GetPC() computer {
	return b.pc
}

func NewGamingComputer(a computer) PCBuilder {
	return &gamingComputer{pc: a}
}
