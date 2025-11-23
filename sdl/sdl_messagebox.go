package sdl

import (
	"unsafe"

	"github.com/SunshineZzzz/purego-sdl3/internal/convert"
)

type MessageBoxColorType uint32

const (
	MessageBoxColorBackground MessageBoxColorType = iota
	MessageBoxColorText
	MessageBoxColorButtonBorder
	MessageBoxColorButtonBackground
	MessageBoxColorButtonSelected
	MessageBoxColorCount
)

type MessageBoxFlags uint32

const (
	MessageBoxError              MessageBoxFlags = 0x00000010
	MessageBoxWarning            MessageBoxFlags = 0x00000020
	MessageBoxInformation        MessageBoxFlags = 0x00000040
	MessageBoxButtonsLeftToRight MessageBoxFlags = 0x00000080
	MessageBoxButtonsRightToLeft MessageBoxFlags = 0x00000100
)

type MessageBoxButtonFlags uint32

const (
	MessageBoxButtonReturnKeyDefault MessageBoxButtonFlags = 0x00000001
	MessageBoxButtonEscapeKeyDefault MessageBoxButtonFlags = 0x00000002
)

type MessageBoxButtonData struct {
	Flags    MessageBoxButtonFlags
	ButtonID int32
	text     *byte
}

func (m *MessageBoxButtonData) Text() string {
	return convert.ToString(m.text)
}

// SetText sets the text of a button.
func (m *MessageBoxButtonData) SetText(s string) {
	m.text = convert.ToBytePtr(s)
}

type MessageBoxColor struct {
	R, G, B uint8
}

type MessageBoxColorScheme struct {
	Colors [MessageBoxColorCount]MessageBoxColor
}

type MessageBoxData struct {
	Flags       MessageBoxFlags
	Window      *Window
	title       *byte
	message     *byte
	numbuttons  int32
	buttons     *MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme
}

func (m *MessageBoxData) Title() string {
	return convert.ToString(m.title)
}

// SetTitle sets the title of a MessageBox.
func (m *MessageBoxData) SetTitle(s string) {
	m.title = convert.ToBytePtr(s)
}

func (m *MessageBoxData) Message() string {
	return convert.ToString(m.message)
}

// SetMessage sets the message text of a MessageBox.
func (m *MessageBoxData) SetMessage(s string) {
	m.message = convert.ToBytePtr(s)
}

func (m *MessageBoxData) Buttons() []MessageBoxButtonData {
	return unsafe.Slice(m.buttons, m.numbuttons)
}

func (m *MessageBoxData) SetButtons(buttons ...MessageBoxButtonData) {
	m.numbuttons = int32(len(buttons))
	if m.numbuttons == 0 {
		m.buttons = nil
		return
	}
	m.buttons = &buttons[0]
}

func ShowMessageBox(messageboxdata *MessageBoxData, buttonid *int32) bool {
	return sdlShowMessageBox(messageboxdata, buttonid)
}

func ShowSimpleMessageBox(flags MessageBoxFlags, title string, message string, window *Window) bool {
	return sdlShowSimpleMessageBox(flags, title, message, window)
}
