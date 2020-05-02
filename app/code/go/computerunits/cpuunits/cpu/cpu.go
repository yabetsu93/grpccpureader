package cpuunits

import (
	"grpccpureader/app/code/go/generator"
	"grpccpureader/protos/pb/go/cpu"
)

// NewCPU returns new cpu
func NewCPU() *cpu.CPU {
	brand := generator.RandomCPUBrand()
	name := generator.RandomCPUName(brand)
	numberCores := generator.RandomInt(2, 8)
	numberThreads := generator.RandomInt(numberCores, 12)
	minGHZ := generator.RandomFloat64(2.0, 3.5)
	maxGHZ := generator.RandomFloat64(minGHZ, 5.0)

	cpu := &cpu.CPU{
		ProcessUnitInfo: &cpu.ProcessUnitInfo{
			Brand:  brand,
			Name:   name,
			MinGhz: minGHZ,
			MaxGhz: maxGHZ,
		},
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
	}

	return cpu
}
