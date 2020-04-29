package storageunits

import (
	"grpccpureader/app/code/go/generator"
	"grpccpureader/protos/pb/go/memory"
	"grpccpureader/protos/pb/go/storage"
)

// NewHDD returns a new hdd
func NewHDD() *storage.Storage {
	hdd := &storage.Storage{
		Driver: storage.Storage_HDD,
		Memory: &memory.Memory{
			Value: uint64(generator.RandomInt(128, 1024)),
			Unit:  memory.Memory_GIGABYTE,
		},
	}

	return hdd
}
