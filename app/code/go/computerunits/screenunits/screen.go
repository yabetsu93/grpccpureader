package screenunits

import (
	"grpccpureader/app/code/go/generator"
	"grpccpureader/protos/pb/go/screen"
)

func NewScreen() *screen.Screen {
	screen := &screen.Screen{
		SizeInch:   generator.RandomFloat32(1080, 4320),
		Resolution: generator.RandomResolution(),
		Panel:      generator.RandomPanel(),
		Multitouch: generator.RandomBool(),
	}

	return screen
}
