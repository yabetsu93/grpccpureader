package storageunits

import (
	"grpccpureader/app/code/go/generator"
	"grpccpureader/protos/pb/go/memory"
	"grpccpureader/protos/pb/go/storage"
)

// NewHDD returns a new hdd
func NewHDD() *storage.Storage {
	memory := &memory.Memory{
		Value: uint64(generator.RandomInt(128, 1064)),
		Unit:  memory.Memory_TERABYTE,
	}

	hdd := &storage.Storage{
		Driver: storage.Storage_HDD,
		Memory: &memory,
	}

	return hdd
}
