package cpuunits

import (
	"grpccpureader/app/code/go/generator"
	"grpccpureader/protos/pb/go/cpu"
	"grpccpureader/protos/pb/go/memory"
)

// NewGPU returns new gpu
func NewGPU() *cpu.GPU {
	brand := generator.RandomCPUBrand()
	name := generator.RandomCPUName(brand)
	minGHZ := generator.RandomFloat64(1.0, 1.5)
	maxGHZ := generator.RandomFloat64(minGHZ, 2.0)
	memory := &memory.Memory{
		Value: uint64(generator.RandomInt(2, 6)),
		Unit:  memory.Memory_GIGABYTE,
	}

	gpu := &cpu.GPU{
		ProcessUnitInfo: &cpu.ProcessUnitInfo{
			Brand:  brand,
			Name:   name,
			MinGhz: minGHZ,
			MaxGhz: maxGHZ,
		},
		Memory: memory,
	}

	return gpu
}
