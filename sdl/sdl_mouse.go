package sdl

import (
	"unsafe"

	"github.com/SunshineZzzz/purego-sdl3/internal/mem"
)

type MouseWheelDirection uint32

const (
	MouseWheelNormal MouseWheelDirection = iota
	MouseWheelFlipped
)

type SystemCursor uint32

const (
	SystemCursorDefault SystemCursor = iota
	SystemCursorText
	SystemCursorWait
	SystemCursorCrosshair
	SystemCursorProgress
	SystemCursorNWSEResize
	SystemCursorNESWResize
	SystemCursorEWResize
	SystemCursorNSResize
	SystemCursorMove
	SystemCursorNotAllowed
	SystemCursorPointer
	SystemCursorNWResize
	SystemCursorNResize
	SystemCursorNEResize
	SystemCursorEResize
	SystemCursorSEResize
	SystemCursorSResize
	SystemCursorSWResize
	SystemCursorWResize
	SystemCursorCount
)

type MouseID uint32

type MouseButtonFlags uint32

const (
	ButtonLeft MouseButtonFlags = iota + 1
	ButtonMiddle
	ButtonRight
	ButtonX1
	ButtonX2

	ButtonLMask  MouseButtonFlags = 1 << (ButtonLeft - 1)
	ButtonMMask  MouseButtonFlags = 1 << (ButtonMiddle - 1)
	ButtonRMask  MouseButtonFlags = 1 << (ButtonRight - 1)
	ButtonX1Mask MouseButtonFlags = 1 << (ButtonX1 - 1)
	ButtonX2Mask MouseButtonFlags = 1 << (ButtonX2 - 1)
)

type Cursor struct{}

func CaptureMouse(enabled bool) bool {
	return sdlCaptureMouse(enabled)
}

func CreateColorCursor(surface *Surface, hotX int32, hotY int32) *Cursor {
	return sdlCreateColorCursor(surface, hotX, hotY)
}

func CreateCursor(data *uint8, mask *uint8, w int32, h int32, hotX int32, hotY int32) *Cursor {
	return sdlCreateCursor(data, mask, w, h, hotX, hotY)
}

func CreateSystemCursor(id SystemCursor) *Cursor {
	return sdlCreateSystemCursor(id)
}

// CursorVisible returns true if the cursor is being shown, or false if the cursor is hidden.
func CursorVisible() bool {
	return sdlCursorVisible()
}

func DestroyCursor(cursor *Cursor) {
	sdlDestroyCursor(cursor)
}

func GetCursor() *Cursor {
	return sdlGetCursor()
}

func GetDefaultCursor() *Cursor {
	return sdlGetDefaultCursor()
}

func GetGlobalMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetGlobalMouseState(x, y)
}

// GetMice returns a list of currently connected mice or nil on failure.
func GetMice() []MouseID {
	var count int32
	mice := sdlGetMice(&count)
	defer Free(unsafe.Pointer(mice))
	return mem.Copy(mice, count)
}

// GetMouseFocus returns the window with mouse focus.
func GetMouseFocus() *Window {
	return sdlGetMouseFocus()
}

// GetMouseNameForID returns the name of the selected mouse.
//
// On failure or if the mouse doesn't have a name, this function returns "".
func GetMouseNameForID(instanceId MouseID) string {
	return sdlGetMouseNameForID(instanceId)
}

// GetMouseState queries SDL's cache for the synchronous mouse button state and the window-relative SDL-cursor position.
func GetMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetMouseState(x, y)
}

func GetRelativeMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetRelativeMouseState(x, y)
}

func GetWindowRelativeMouseMode(window *Window) bool {
	return sdlGetWindowRelativeMouseMode(window)
}

// HasMouse returns true if a mouse is connected, false otherwise.
func HasMouse() bool {
	return sdlHasMouse()
}

// HideCursor hides the cursor.
func HideCursor() bool {
	return sdlHideCursor()
}

func SetCursor(cursor *Cursor) bool {
	return sdlSetCursor(cursor)
}

func SetWindowRelativeMouseMode(window *Window, enabled bool) bool {
	return sdlSetWindowRelativeMouseMode(window, enabled)
}

// ShowCursor shows the cursor.
func ShowCursor() bool {
	return sdlShowCursor()
}

func WarpMouseGlobal(x float32, y float32) bool {
	return sdlWarpMouseGlobal(x, y)
}

// WarpMouseInWindow moves the mouse cursor to the given position within the window.
func WarpMouseInWindow(window *Window, x float32, y float32) {
	sdlWarpMouseInWindow(window, x, y)
}
