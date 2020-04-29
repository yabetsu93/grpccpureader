package generator



func NewKeyboard() *keyboard. {
	keyboard := &keyboard.Keyboard{
		Layout:  randomKeyboardLayout,
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *cpu.CPU {
	cpu.
}
