package generator

import (
	pb "grpccpureader/protos/pb/keyboard"
	"math/rand"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWETZ
	default:
		return pb.Keyboard_AZERT
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
