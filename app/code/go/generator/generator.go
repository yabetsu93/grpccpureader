package generator

import (
	"grpccpureader/protos/pb/go/cpu"
	"grpccpureader/protos/pb/go/keyboard"
	"grpccpureader/protos/pb/go/memory"
	"grpccpureader/protos/pb/go/storage"
)

// NewKeyboard returns new keyboad
func NewKeyboard() *keyboard.Keyboard {
	keyboard := &keyboard.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns new cpu
func NewCPU() *cpu.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGHZ := randomFloat64(2.0, 3.5)
	maxGHZ := randomFloat64(minGHZ, 5.0)

	cpu := &cpu.CPU{
		ProcessUnitInfo: *cpu.ProcessUnitInfo{
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

// NewGPU returns new gpu
func NewGPU() *cpu.GPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	minGHZ := randomFloat64(1.0, 1.5)
	maxGHZ := randomFloat64(minGHZ, 2.0)
	memory := &memory.Memory{
		Value: uint64(randomInt(2, 6)),
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

// NewRam returns a new ram
func NewRam() *memory.Memory {
	ram := &memory.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  memory.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new ssd
func NewSSD() *storage.Storage {
	ssd := &storage.Storage{
		Driver: storage.Storage_SSD,
		Memory: &memory.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  memory.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new hdd
func NewHDD() *storage.Storage {
	hdd := &storage.Storage{
		Driver: storage.Storage_HDD,
		Memory: &memory.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  memory.Memory_GIGABYTE,
		},
	}

	return hdd
}
