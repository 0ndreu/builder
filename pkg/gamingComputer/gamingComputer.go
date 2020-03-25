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
	b.pc.CPU("i7")
	return b
}

func (b *gamingComputer) SetRAM() PCBuilder {
	b.pc.Ram(32)
	return b
}

func (b *gamingComputer) SetHDD() PCBuilder {
	b.pc.HDD(1024)
	return b
}

func (b *gamingComputer) GetPC() computer {
	return b.pc
}

func NewGamingComputer(a computer) PCBuilder {
	return &gamingComputer{pc: a}
}
