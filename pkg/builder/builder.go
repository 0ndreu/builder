package builder

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

type Director interface {
	ConstructPC()
}

type director struct {
	builder PCBuilder
}

func (d *director) ConstructPC() {
	d.builder.SetCPU()
	d.builder.SetHDD()
	d.builder.SetRAM()
	d.builder.GetPC()
}

func NewDirector(p PCBuilder) Director {
	return &director{
		builder: p,
	}
}
