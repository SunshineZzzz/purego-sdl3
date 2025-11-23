package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/SunshineZzzz/purego-sdl3/internal/convert"
)

type EventAction uint32

const (
	AddEvent EventAction = iota
	PeekEvent
	GetEvent
)

type EventType uint32

const (
	EventFirst                      EventType = 0x0
	EventQuit                       EventType = 0x100
	EventTerminating                EventType = 0x101
	EventLowMemory                  EventType = 0x102
	EventWillEnterBackground        EventType = 0x103
	EventDidEnterBackground         EventType = 0x104
	EventWillEnterForeground        EventType = 0x105
	EventDidEnterForeground         EventType = 0x106
	EventLocaleChanged              EventType = 0x107
	EventSystemThemeChanged         EventType = 0x108
	EventDisplayOrientation         EventType = 0x151
	EventDisplayFirst               EventType = 0x151
	EventDisplayAdded               EventType = 0x152
	EventDisplayRemoved             EventType = 0x153
	EventDisplayMoved               EventType = 0x154
	EventDisplayDesktopModeChanged  EventType = 0x155
	EventDisplayCurrentModeChanged  EventType = 0x156
	EventDisplayContentScaleChanged EventType = 0x157
	EventDisplayLast                EventType = 0x157
	EventWindowShown                EventType = 0x202
	EventWindowFirst                EventType = 0x202
	EventWindowHidden               EventType = 0x203
	EventWindowExposed              EventType = 0x204
	EventWindowMoved                EventType = 0x205
	EventWindowResized              EventType = 0x206
	EventWindowPixelSizeChanged     EventType = 0x207
	EventWindowMetalViewResized     EventType = 0x208
	EventWindowMinimized            EventType = 0x209
	EventWindowMaximized            EventType = 0x20A
	EventWindowRestored             EventType = 0x20B
	EventWindowMouseEnter           EventType = 0x20C
	EventWindowMouseLeave           EventType = 0x20D
	EventWindowFocusGained          EventType = 0x20E
	EventWindowFocusLost            EventType = 0x20F
	EventWindowCloseRequested       EventType = 0x210
	EventWindowHitTest              EventType = 0x211
	EventWindowICCProfChanged       EventType = 0x212
	EventWindowDisplayChanged       EventType = 0x213
	EventWindowDisplayScaleChanged  EventType = 0x214
	EventWindowSafeAreaChanged      EventType = 0x215
	EventWindowOccluded             EventType = 0x216
	EventWindowEnterFullscreen      EventType = 0x217
	EventWindowLeaveFullscreen      EventType = 0x218
	EventWindowDestroyed            EventType = 0x219
	EventWindowHDRStateChanged      EventType = 0x21A
	EventWindowLast                 EventType = 0x21A
	EventKeyDown                    EventType = 0x300
	EventKeyUp                      EventType = 0x301
	EventTextEditing                EventType = 0x302
	EventTextInput                  EventType = 0x303
	EventKeymapChanged              EventType = 0x304
	EventKeyboardAdded              EventType = 0x305
	EventKeyboardRemoved            EventType = 0x306
	EventTextEditingCandidates      EventType = 0x307
	EventMouseMotion                EventType = 0x400
	EventMouseButtonDown            EventType = 0x401
	EventMouseButtonUp              EventType = 0x402
	EventMouseWheel                 EventType = 0x403
	EventMouseAdded                 EventType = 0x404
	EventMouseRemoved               EventType = 0x405
	EventJoystickAxisMotion         EventType = 0x600
	EventJoystickBallMotion         EventType = 0x601
	EventJoystickHatMotion          EventType = 0x602
	EventJoystickButtonDown         EventType = 0x603
	EventJoystickButtonUp           EventType = 0x604
	EventJoystickAdded              EventType = 0x605
	EventJoystickRemoved            EventType = 0x606
	EventJoystickBatteryUpdated     EventType = 0x607
	EventJoystickUpdateComplete     EventType = 0x608
	EventGamepadAxisMotion          EventType = 0x650
	EventGamepadButtonDown          EventType = 0x651
	EventGamepadButtonUp            EventType = 0x652
	EventGamepadAdded               EventType = 0x653
	EventGamepadRemoved             EventType = 0x654
	EventGamepadRemapped            EventType = 0x655
	EventGamepadTouchpadDown        EventType = 0x656
	EventGamepadTouchpadMotion      EventType = 0x657
	EventGamepadTouchpadUp          EventType = 0x658
	EventGamepadSensorUpdate        EventType = 0x659
	EventGamepadUpdateComplete      EventType = 0x65A
	EventGamepadSteamHandleUpdated  EventType = 0x65B
	EventFingerDown                 EventType = 0x700
	EventFingerUp                   EventType = 0x701
	EventFingerMotion               EventType = 0x702
	EventFingerCanceled             EventType = 0x703
	EventClipboardUpdate            EventType = 0x900
	EventDropFile                   EventType = 0x1000
	EventDropText                   EventType = 0x1001
	EventDropBegin                  EventType = 0x1002
	EventDropComplete               EventType = 0x1003
	EventDropPosition               EventType = 0x1004
	EventAudioDeviceAdded           EventType = 0x1100
	EventAudioDeviceRemoved         EventType = 0x1101
	EventAudioDeviceFormatChanged   EventType = 0x1102
	EventSensorUpdate               EventType = 0x1200
	EventPenProximityIn             EventType = 0x1300
	EventPenProximityOut            EventType = 0x1301
	EventPenDown                    EventType = 0x1302
	EventPenUp                      EventType = 0x1303
	EventPenButtonDown              EventType = 0x1304
	EventPenButtonUp                EventType = 0x1305
	EventPenMotion                  EventType = 0x1306
	EventPenAxis                    EventType = 0x1307
	EventCameraDeviceAdded          EventType = 0x1400
	EventCameraDeviceRemoved        EventType = 0x1401
	EventCameraDeviceApproved       EventType = 0x1402
	EventCameraDeviceDenied         EventType = 0x1403
	EventRenderTargetsReset         EventType = 0x2000
	EventRenderDeviceReset          EventType = 0x2001
	EventRenderDeviceLost           EventType = 0x2002
	EventPrivate0                   EventType = 0x4000
	EventPrivate1                   EventType = 0x4001
	EventPrivate2                   EventType = 0x4002
	EventPrivate3                   EventType = 0x4003
	EventPollSentinel               EventType = 0x7F00
	EventUser                       EventType = 0x8000
	EventLast                       EventType = 0xFFFF
	EventEnumPadding                EventType = 0x7FFFFFFF
)

type Event [128]byte

func (e *Event) Type() EventType {
	return *(*EventType)(unsafe.Pointer(e))
}

func (e *Event) Common() CommonEvent {
	return *(*CommonEvent)(unsafe.Pointer(e))
}

func (e *Event) Display() DisplayEvent {
	return *(*DisplayEvent)(unsafe.Pointer(e))
}

func (e *Event) Window() WindowEvent {
	return *(*WindowEvent)(unsafe.Pointer(e))
}

func (e *Event) KDevice() KeyboardDeviceEvent {
	return *(*KeyboardDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) Key() KeyboardEvent {
	return *(*KeyboardEvent)(unsafe.Pointer(e))
}

func (e *Event) Edit() TextEditingEvent {
	return *(*TextEditingEvent)(unsafe.Pointer(e))
}

func (e *Event) EditCandidates() TextEditingCandidatesEvent {
	return *(*TextEditingCandidatesEvent)(unsafe.Pointer(e))
}

func (e *Event) Text() TextInputEvent {
	return *(*TextInputEvent)(unsafe.Pointer(e))
}

func (e *Event) MDevice() MouseDeviceEvent {
	return *(*MouseDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) Motion() MouseMotionEvent {
	return *(*MouseMotionEvent)(unsafe.Pointer(e))
}

func (e *Event) Button() MouseButtonEvent {
	return *(*MouseButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) Wheel() MouseWheelEvent {
	return *(*MouseWheelEvent)(unsafe.Pointer(e))
}

func (e *Event) JDevice() JoyDeviceEvent {
	return *(*JoyDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) JAxis() JoyAxisEvent {
	return *(*JoyAxisEvent)(unsafe.Pointer(e))
}

func (e *Event) JBall() JoyBallEvent {
	return *(*JoyBallEvent)(unsafe.Pointer(e))
}

func (e *Event) JHat() JoyHatEvent {
	return *(*JoyHatEvent)(unsafe.Pointer(e))
}

func (e *Event) JButton() JoyButtonEvent {
	return *(*JoyButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) JBattery() JoyBatteryEvent {
	return *(*JoyBatteryEvent)(unsafe.Pointer(e))
}

func (e *Event) GDevice() GamepadDeviceEvent {
	return *(*GamepadDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) GAxis() GamepadAxisEvent {
	return *(*GamepadAxisEvent)(unsafe.Pointer(e))
}

func (e *Event) GButton() GamepadButtonEvent {
	return *(*GamepadButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) GTouchpad() GamepadTouchpadEvent {
	return *(*GamepadTouchpadEvent)(unsafe.Pointer(e))
}

func (e *Event) GSensor() GamepadSensorEvent {
	return *(*GamepadSensorEvent)(unsafe.Pointer(e))
}

func (e *Event) ADevice() AudioDeviceEvent {
	return *(*AudioDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) CDevice() CameraDeviceEvent {
	return *(*CameraDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) Sensor() SensorEvent {
	return *(*SensorEvent)(unsafe.Pointer(e))
}

func (e *Event) Quit() QuitEvent {
	return *(*QuitEvent)(unsafe.Pointer(e))
}

func (e *Event) User() UserEvent {
	return *(*UserEvent)(unsafe.Pointer(e))
}

func (e *Event) TFinger() TouchFingerEvent {
	return *(*TouchFingerEvent)(unsafe.Pointer(e))
}
func (e *Event) PProximity() PenProximityEvent {
	return *(*PenProximityEvent)(unsafe.Pointer(e))
}
func (e *Event) PTouch() PenTouchEvent {
	return *(*PenTouchEvent)(unsafe.Pointer(e))
}
func (e *Event) PMotion() PenMotionEvent {
	return *(*PenMotionEvent)(unsafe.Pointer(e))
}
func (e *Event) PButton() PenButtonEvent {
	return *(*PenButtonEvent)(unsafe.Pointer(e))
}
func (e *Event) PAxis() PenAxisEvent {
	return *(*PenAxisEvent)(unsafe.Pointer(e))
}
func (e *Event) Render() RenderEvent {
	return *(*RenderEvent)(unsafe.Pointer(e))
}

func (e *Event) Drop() DropEvent {
	return *(*DropEvent)(unsafe.Pointer(e))
}

func (e *Event) Clipboard() ClipboardEvent {
	return *(*ClipboardEvent)(unsafe.Pointer(e))
}

// CommonEvent fields are shared by every event.
type CommonEvent struct {
	Type     EventType
	Reserved uint32
	// Timestamp in nanoseconds, populated using [GetTicksNS].
	Timestamp uint64
}

type DisplayEvent struct {
	CommonEvent
	DisplayID DisplayID
	Data1     int32
	Data2     int32
}

type WindowEvent struct {
	CommonEvent
	WindowID WindowID
	Data1    int32
	Data2    int32
}

type KeyboardDeviceEvent struct {
	CommonEvent
	Which KeyboardID
}

type KeyboardEvent struct {
	CommonEvent
	WindowID WindowID
	Which    KeyboardID
	Scancode Scancode
	Key      Keycode
	Mod      Keymod
	Raw      uint16
	Down     bool
	Repeat   bool
}

type TextEditingEvent struct {
	CommonEvent
	WindowID WindowID
	text     *byte
	Start    int32
	Length   int32
}

func (t *TextEditingEvent) Text() string {
	return convert.ToString(t.text)
}

type TextEditingCandidatesEvent struct {
	CommonEvent
	WindowID          WindowID
	candidates        **byte
	numCandidates     int32
	SelectedCandidate int32
	Horizontal        bool
	Padding1          uint8
	Padding2          uint8
	Padding3          uint8
}

// Candidates returns the list of candidates,
// or empty if there are no candidates available.
func (t *TextEditingCandidatesEvent) Candidates() []string {
	if t.candidates == nil || t.numCandidates <= 0 {
		return []string{}
	}
	candidates := make([]string, t.numCandidates)
	for i, v := range unsafe.Slice(t.candidates, t.numCandidates) {
		candidates[i] = convert.ToString(v)
	}
	return candidates
}

type TextInputEvent struct {
	CommonEvent
	WindowID WindowID
	text     *byte
}

func (t *TextInputEvent) Text() string {
	return convert.ToString(t.text)
}

type MouseDeviceEvent struct {
	CommonEvent
	Which MouseID
}

type MouseMotionEvent struct {
	CommonEvent
	WindowID WindowID
	Which    MouseID
	State    MouseButtonFlags
	X        float32
	Y        float32
	Xrel     float32
	Yrel     float32
}

type MouseButtonEvent struct {
	CommonEvent
	WindowID WindowID
	Which    MouseID
	Button   uint8
	Down     bool
	Clicks   uint8
	Padding  uint8
	X        float32
	Y        float32
}

type MouseWheelEvent struct {
	CommonEvent
	WindowID  WindowID
	Which     MouseID
	X         float32
	Y         float32
	Direction MouseWheelDirection
	MouseX    float32
	MouseY    float32
}

type JoyAxisEvent struct {
	CommonEvent
	Which    JoystickID
	Axis     uint8
	Padding1 uint8
	Padding2 uint8
	Padding3 uint8
	Value    int16
	Padding4 uint16
}

type JoyBallEvent struct {
	CommonEvent
	Which    JoystickID
	Ball     uint8
	Padding1 uint8
	Padding2 uint8
	Padding3 uint8
	Xrel     int16
	Yrel     int16
}

type JoyHatEvent struct {
	CommonEvent
	Which    JoystickID
	Hat      uint8
	Value    uint8
	Padding1 uint8
	Padding2 uint8
}

type JoyButtonEvent struct {
	CommonEvent
	Which    JoystickID
	Button   uint8
	Down     bool
	Padding1 uint8
	Padding2 uint8
}

type JoyDeviceEvent struct {
	CommonEvent
	Which JoystickID
}

type JoyBatteryEvent struct {
	CommonEvent
	Which   JoystickID
	State   PowerState
	Percent int32
}

type GamepadAxisEvent struct {
	CommonEvent
	Which    JoystickID
	Axis     uint8
	Padding1 uint8
	Padding2 uint8
	Padding3 uint8
	Value    int16
	Padding4 uint16
}

type GamepadButtonEvent struct {
	CommonEvent
	Which    JoystickID
	Button   uint8
	Down     bool
	Padding1 uint8
	Padding2 uint8
}

type GamepadDeviceEvent struct {
	CommonEvent
	Which JoystickID
}

type GamepadTouchpadEvent struct {
	CommonEvent
	Which    JoystickID
	Touchpad int32
	Finger   int32
	X        float32
	Y        float32
	Pressure float32
}

type GamepadSensorEvent struct {
	CommonEvent
	Which           JoystickID
	Sensor          int32
	Data            [3]float32
	SensorTimestamp uint64
}

type AudioDeviceEvent struct {
	CommonEvent
	Which     AudioDeviceID
	Recording bool
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
}

type CameraDeviceEvent struct {
	CommonEvent
	Which CameraID
}

type RenderEvent struct {
	CommonEvent
	WindowID WindowID
}

type TouchFingerEvent struct {
	CommonEvent
	TouchID  TouchID
	FingerID FingerID
	X        float32
	Y        float32
	Dx       float32
	Dy       float32
	Pressure float32
	WindowID WindowID
}

type PenProximityEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
}

type PenMotionEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X        float32
	Y        float32
}

type PenTouchEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X        float32
	Y        float32
	Eraser   bool
	Down     bool
}

type PenButtonEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X        float32
	Y        float32
	Button   uint8
	Down     bool
}

type PenAxisEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X        float32
	Y        float32
	Axis     PenAxis
	Value    float32
}

type DropEvent struct {
	CommonEvent
	WindowID WindowID
	X        float32
	Y        float32
	source   *byte
	data     *byte
}

func (d *DropEvent) Source() string {
	return convert.ToString(d.source)
}

func (d *DropEvent) Data() string {
	return convert.ToString(d.data)
}

type ClipboardEvent struct {
	CommonEvent
	Owner        bool
	NumMimeTypes int32
	mimeTypes    **byte
}

func (c *ClipboardEvent) MimeTypes() []string {
	mimeTypes := make([]string, c.NumMimeTypes)
	for i, v := range unsafe.Slice(c.mimeTypes, c.NumMimeTypes) {
		mimeTypes[i] = convert.ToString(v)
	}

	return mimeTypes
}

type SensorEvent struct {
	CommonEvent
	Which           SensorID
	Data            [6]float32
	SensorTimestamp uint64
}

type QuitEvent struct {
	CommonEvent
}

type UserEvent struct {
	CommonEvent
	WindowID WindowID
	Code     int32
	Data1    unsafe.Pointer
	Data2    unsafe.Pointer
}

// EventFilter is a C function pointer used for callbacks that watch the event queue. Use [NewEventFilter] for creation.
type EventFilter uintptr

// NewEventFilter converts the Go function to a C function pointer.
func NewEventFilter(filter func(userdata unsafe.Pointer, event *Event) bool) EventFilter {
	// workaround to avoid panic "expected function with one uintptr-sized result" on Windows
	cb := purego.NewCallback(func(userdata unsafe.Pointer, event *Event) uintptr {
		if filter(userdata, event) {
			return 1
		}
		return 0
	})

	return EventFilter(cb)
}

// PollEvent polls for currently pending events.
func PollEvent(event *Event) bool {
	ret, _, _ := purego.SyscallN(sdlPollEvent, uintptr(unsafe.Pointer(event)))
	return byte(ret) != 0
}

// AddEventWatch adds a callback to be triggered when an event is added to the event queue.
func AddEventWatch(filter EventFilter, userdata unsafe.Pointer) bool {
	return sdlAddEventWatch(filter, userdata)
}

// EventEnabled returns true if the event is being processed, false otherwise.
func EventEnabled(eventType EventType) bool {
	return sdlEventEnabled(eventType)
}

func FilterEvents(filter EventFilter, userdata unsafe.Pointer) {
	sdlFilterEvents(filter, userdata)
}

// FlushEvent clears events of a specific type from the event queue.
func FlushEvent(eventType EventType) {
	sdlFlushEvent(eventType)
}

// FlushEvents clears events of a range of types from the event queue.
func FlushEvents(minType, maxType EventType) {
	sdlFlushEvents(minType, maxType)
}

// GetEventFilter queries the current event filter.
func GetEventFilter(filter *EventFilter, userdata *unsafe.Pointer) bool {
	return sdlGetEventFilter(filter, userdata)
}

// GetWindowFromEvent returns the associated window with an event or nil if there is none.
func GetWindowFromEvent(event *Event) *Window {
	return sdlGetWindowFromEvent(event)
}

// HasEvent checks for the existence of a certain event type in the event queue.
func HasEvent(eventType EventType) bool {
	return sdlHasEvent(eventType)
}

// HasEvents checks for the existence of certain event types in the event queue.
//
// Returns true if events with type >= minType and <= maxType are present, or false if not.
func HasEvents(minType, maxType EventType) bool {
	return sdlHasEvents(minType, maxType)
}

// PeepEvents checks the event queue for messages and optionally returns them.
//
// Example:
//
//	sdl.PumpEvents()
//	var events [2]sdl.Event
//	sdl.PeepEvents(&events[0], 2, sdl.PeekEvent, sdl.EventFirst, sdl.EventLast)
func PeepEvents(events *Event, numevents int32, action EventAction, minType, maxType EventType) int32 {
	return sdlPeepEvents(events, numevents, action, minType, maxType)
}

// PumpEvents updates the event queue and internal input device state.
func PumpEvents() {
	sdlPumpEvents()
}

// PushEvent adds an event to the event queue.
func PushEvent(event *Event) bool {
	return sdlPushEvent(event)
}

// RegisterEvents allocates a set of user-defined events, and return the beginning event number for that set of events.
func RegisterEvents(numevents int32) uint32 {
	return sdlRegisterEvents(numevents)
}

// RemoveEventWatch removes an event watch callback added with [AddEventWatch].
//
// This function takes the same input as [AddEventWatch] to identify and delete the corresponding callback.
func RemoveEventWatch(filter EventFilter, userdata unsafe.Pointer) {
	sdlRemoveEventWatch(filter, userdata)
}

// SetEventEnabled sets the state of processing events by type.
func SetEventEnabled(eventType EventType, enabled bool) {
	sdlSetEventEnabled(eventType, enabled)
}

// SetEventFilter sets up a filter to process all events before they are added to the internal event queue.
func SetEventFilter(filter EventFilter, userdata unsafe.Pointer) {
	sdlSetEventFilter(filter, userdata)
}

// WaitEvent waits indefinitely for the next available event.
func WaitEvent(event *Event) bool {
	return sdlWaitEvent(event)
}

// WaitEventTimeout waits until the specified timeout (in milliseconds) for the next available event.
func WaitEventTimeout(event *Event, timeoutMS int32) bool {
	return sdlWaitEventTimeout(event, timeoutMS)
}
