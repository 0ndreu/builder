package builder

import "builder/pkg/myComputer"

type PCBuilder interface {
	SetCPU() PCBuilder
	SetRAM() PCBuilder
	SetHDD() PCBuilder
	GetPC() myComputer.MyComputer
}
