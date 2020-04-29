package memoryunits

import (
	"grpccpureader/app/code/go/generator"
	"grpccpureader/protos/pb/go/memory"
)

// NewRam returns a new ram
func NewRam() *memory.Memory {
	ram := &memory.Memory{
		Value: uint64(generator.RandomInt(2, 6)),
		Unit:  memory.Memory_GIGABYTE,
	}

	return ram
}
