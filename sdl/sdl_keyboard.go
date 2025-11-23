package sdl

import (
	"unsafe"

	"github.com/SunshineZzzz/purego-sdl3/internal/mem"
)

const (
	PropTextInputTypeNumber             = "SDL.textinput.type"
	PropTextInputCapitalizationNumber   = "SDL.textinput.capitalization"
	PropTextInputAutocorrectBoolean     = "SDL.textinput.autocorrect"
	PropTextInputMultilineBoolean       = "SDL.textinput.multiline"
	PropTextInputAndroidInputTypeNumber = "SDL.textinput.android.inputtype"
)

type Capitalization uint32

const (
	CapitalizeNone Capitalization = iota
	CapitalizeSentences
	CapitalizeWords
	CapitalizeLetters
)

type TextInputType uint32

const (
	TextInputTypeText TextInputType = iota
	TextInputTypeTextName
	TextInputTypeTextEmail
	TextInputTypeTextUsername
	TextInputTypeTextPasswordHidden
	TextInputTypeTextPasswordVisible
	TextInputTypeNumber
	TextInputTypeNumberPasswordHidden
	TextInputTypeNumberPasswordVisible
)

type KeyboardID uint32

func ClearComposition(window *Window) bool {
	return sdlClearComposition(window)
}

func GetKeyboardFocus() *Window {
	return sdlGetKeyboardFocus()
}

func GetKeyboardNameForID(instanceId KeyboardID) string {
	return sdlGetKeyboardNameForID(instanceId)
}

func GetKeyboards() []KeyboardID {
	var count int32
	keyboards := sdlGetKeyboards(&count)
	defer Free(unsafe.Pointer(keyboards))
	return mem.Copy(keyboards, count)
}

func GetKeyboardState() []bool {
	var numkeys int32
	state := sdlGetKeyboardState(&numkeys)
	return unsafe.Slice(state, numkeys)
}

func GetKeyFromName(name string) Keycode {
	return sdlGetKeyFromName(name)
}

func GetKeyFromScancode(scancode Scancode, modstate Keymod, keyEvent bool) Keycode {
	return sdlGetKeyFromScancode(scancode, modstate, keyEvent)
}

// GetKeyName returns a human-readable name for a key.
func GetKeyName(key Keycode) string {
	return sdlGetKeyName(key)
}

// GetModState returns an OR'd combination of the modifier keys for the keyboard.
func GetModState() Keymod {
	return sdlGetModState()
}

func GetScancodeFromKey(key Keycode, modstate *Keymod) Scancode {
	return sdlGetScancodeFromKey(key, modstate)
}

func GetScancodeFromName(name string) Scancode {
	return sdlGetScancodeFromName(name)
}

func GetScancodeName(scancode Scancode) string {
	return sdlGetScancodeName(scancode)
}

func GetTextInputArea(window *Window, rect *Rect, cursor *int32) bool {
	return sdlGetTextInputArea(window, rect, cursor)
}

func HasKeyboard() bool {
	return sdlHasKeyboard()
}

func HasScreenKeyboardSupport() bool {
	return sdlHasScreenKeyboardSupport()
}

func ResetKeyboard() {
	sdlResetKeyboard()
}

func ScreenKeyboardShown(window *Window) bool {
	return sdlScreenKeyboardShown(window)
}

func SetModState(modstate Keymod) {
	sdlSetModState(modstate)
}

func SetScancodeName(scancode Scancode, name string) bool {
	return sdlSetScancodeName(scancode, name)
}

func SetTextInputArea(window *Window, rect *Rect, cursor int32) bool {
	return sdlSetTextInputArea(window, rect, cursor)
}

// StartTextInput starts accepting Unicode text input events in a window.
func StartTextInput(window *Window) bool {
	return sdlStartTextInput(window)
}

func StartTextInputWithProperties(window *Window, props PropertiesID) bool {
	return sdlStartTextInputWithProperties(window, props)
}

// StopTextInput stops receiving any text input events in a window.
func StopTextInput(window *Window) bool {
	return sdlStopTextInput(window)
}

func TextInputActive(window *Window) bool {
	return sdlTextInputActive(window)
}
