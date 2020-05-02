package laptopunits

import (
	"grpccpureader/app/code/go/generator"
	"grpccpureader/protos/pb/go/cpu"
	"grpccpureader/protos/pb/go/laptop"
)

func NewLaptop() *laptop.Laptop {
	brand := generator.RandomLaptopBrand()
	name := generator.RandomLaptopName(brand)
	laptop := &laptop.Laptop{
		Id:     generator.RandomID(),
		Name:   name,
		Brand:  brand,
		Cpu:    generator.RandomCPU(),
		Gpus:   []*cpu.GPU{RandomGPU()},
		Ram:    generator.RandomMemory(),
		Screen: generator.RandomScreen(),
	}

	return laptop
}
