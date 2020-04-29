package keyboardunits

import (
	"grpccpureader/app/code/go/generator"
	"grpccpureader/protos/pb/go/keyboard"
)

// NewKeyboard returns new keyboad
func NewKeyboard() *keyboard.Keyboard {
	keyboard := &keyboard.Keyboard{
		Layout:  generator.RandomKeyboardLayout(),
		Backlit: generator.RandomBool(),
	}
	return keyboard
}
