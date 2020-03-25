package computer

type Computer interface {
	CPU(proc string)
	Ram(size int)
	HDD(size int)
}

type computer struct {
	cpu string
	ram int
	hdd int
}

func (pc *computer) CPU(proc string) {
	pc.cpu = proc
}

func (pc *computer) Ram(size int) {
	pc.ram = size
}

func (pc *computer) HDD(size int) {
	pc.hdd = size
}

func NewComputer() Computer {
	return &computer{
		cpu: "",
		ram: 0,
		hdd: 0,
	}
}
