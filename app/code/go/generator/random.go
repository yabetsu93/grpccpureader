package generator

import (
	"grpccpureader/protos/pb/go/keyboard"
	"grpccpureader/protos/pb/go/memory"
	"grpccpureader/protos/pb/go/screen"
	"math/rand"

	"github.com/google/uuid"
)

// RandomKeyboardLayout returns a random keyboard layout
func RandomKeyboardLayout() keyboard.Keyboard_Layout {
	return randomKbLayout()
}

// RandomBool returns a random bool
func RandomBool() bool {
	return rand.Intn(2) == 1
}

// RandomCPUBrand returns a random cpu brand
func RandomCPUBrand() string {
	return randomStringSet("Intel", "AMD")
}

// RandomStringFromSet returns a random string from set
func RandomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

// RandomCPUName returns a random cpu name
func RandomCPUName(brand string) string {
	return randomCPUName(brand)
}

// RandomInt returns a random integer
func RandomInt(min, max int) int {
	return randomInt(min, max)
}

// RandomFloat32 returns a random float32
func RandomFloat32(min, max float32) float32 {
	return randomFloat32(min, max)
}

// RandomFloat64 returns a random number float64
func RandomFloat64(min, max float64) float64 {
	return randomFloat64(min, max)
}

// RandomPanel returns a random panel
func RandomPanel() screen.Screen_Panel {
	return randomPanel()
}

func RandomResolution() *screen.Screen_Resolution {
	return randomScreenResolution()
}

func randomStringSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomCPUName(brand string) string {
	switch brand {
	case "Intel":
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

func randomKbLayout() keyboard.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return keyboard.Keyboard_QWERTY
	case 2:
		return keyboard.Keyboard_QWETZ
	default:
		return keyboard.Keyboard_AZERT
	}
}

func randomPanel() screen.Screen_Panel {
	if rand.Intn(2) == 1 {
		return screen.Screen_IPS
	}

	return screen.Screen_OLED
}

func randomScreenResolution() *screen.Screen_Resolution {
	height := RandomInt(1080, 4320)
	width := height * 16 / 9

	screen := &screen.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}

	return screen
}

// RandomMemory returns a random memory
func RandomMemory() *memory.Memory {
	memory := &memory.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  memory.Memory_GIGABYTE,
	}

	return memory
}

func RandomID() string {
	return randomID()
}

func randomID() string {
	return uuid.New.String()
}

// RandomLaptopBrand returns a laptop brand
func RandomLaptopBrand() string {
	return randomStringSet("Apple", "Dell")
}

// RandomLaptopName returns a laptop names
func RandomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringSet("Macbook Air", "Macbook Pro")
	default:
		return randomStringSet("Alienware", "Latitude", "Vostro", "XPS")
	}
}
