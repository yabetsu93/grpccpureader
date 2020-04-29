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

func randomCPUBrand() string () {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(a ...string) string {
	n = len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	switch(brand) {
	case 1:
		"Intel": 
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i9-9750H",
			"Core i9-9400F",
		)
	default:
		return randomStringFromSet(
			"Ryzen 7 PRD 2700U",
			"Ryzen 7 PRD 3500U",
			"Ryzen 7 PRD 4300U",
		)
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64() * (max - min)
}
