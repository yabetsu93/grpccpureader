package generator

import (
	pb "grpccpureader/protos/pb/keyboard"
	"math/rand"
)

// RandomKeyboardLayout returns a random keyboard layout
func RandomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWETZ
	default:
		return pb.Keyboard_AZERT
	}
}

// RandomBool returns a random bool
func RandomBool() bool {
	return rand.Intn(2) == 1
}

// RandomCPUBrand returns a random cpu brand
func RandomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

// RandomStringFromSet returns a random string from set
func RandomStringFromSet(a ...string) string {
	return randomStringSet(a)
}

// RandomCPUName returns a random cpu name
func RandomCPUName(brand string) string {
	switch(brand) {
	case 1:
		"Intel": 
		return randomStringSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i9-9750H",
			"Core i9-9400F",
		)
	default:
		return randomStringSet(
			"Ryzen 7 PRD 2700U",
			"Ryzen 7 PRD 3500U",
			"Ryzen 7 PRD 4300U",
		)
	}
}

// RandomInt returns a random integer
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomStringSet(s ...string) string {
	n = len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

// RandomFloat64 returns a random number float64
func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64() * (max - min)
}
