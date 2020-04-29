package storageunits

import (
	"grpccpureader/app/code/go/generator"
	"grpccpureader/protos/pb/go/memory"
	"grpccpureader/protos/pb/go/storage"
)

// NewSSD returns a new ssd
func NewSSD() *storage.Storage {
	ssd := &storage.Storage{
		Driver: storage.Storage_SSD,
		Memory: &memory.Memory{
			Value: uint64(generator.RandomInt(128, 1024)),
			Unit:  memory.Memory_GIGABYTE,
		},
	}

	return ssd
}
