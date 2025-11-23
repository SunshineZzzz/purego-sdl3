package sdl

import (
	"unsafe"

	"github.com/SunshineZzzz/purego-sdl3/internal/mem"
)

type Joystick struct{}

type JoystickID uint32

type JoystickType uint16

const (
	JoystickTypeUnknown JoystickType = iota
	JoystickTypeGamepad
	JoystickTypeWheel
	JoystickTypeArcadeStick
	JoystickTypeFlightStick
	JoystickTypeDancePad
	JoystickTypeGuitar
	JoystickTypeDrumKit
	JoystickTypeArcadePad
	JoystickTypeThrottle
	JoystickTypeCount
)

type JoystickConnectionState int32

const (
	JoystickConnectionInvalid JoystickConnectionState = iota - 1
	JoystickConnectionUnknown
	JoystickConnectionWired
	JoystickConnectionWireless
)

func LockJoysticks() {
	sdlLockJoysticks()
}

func UnlockJoysticks() {
	sdlUnlockJoysticks()
}

func HasJoystick() bool {
	return sdlHasJoystick()
}

func GetJoysticks() []JoystickID {
	var count int32
	joysticks := sdlGetJoysticks(&count)
	if joysticks == nil {
		return nil
	}
	defer Free(unsafe.Pointer(joysticks))
	return mem.Copy(joysticks, count)
}

func GetJoystickNameForID(instanceId JoystickID) string {
	return sdlGetJoystickNameForID(instanceId)
}

func GetJoystickPathForID(instanceId JoystickID) string {
	return sdlGetJoystickPathForID(instanceId)
}

func GetJoystickPlayerIndexForID(instanceId JoystickID) int32 {
	return sdlGetJoystickPlayerIndexForID(instanceId)
}

// func GetJoystickGUIDForID(instanceId JoystickID) GUID {
// 	return sdlGetJoystickGUIDForID(instanceId)
// }

func GetJoystickVendorForID(instanceId JoystickID) uint16 {
	return sdlGetJoystickVendorForID(instanceId)
}

func GetJoystickProductForID(instanceId JoystickID) uint16 {
	return sdlGetJoystickProductForID(instanceId)
}

func GetJoystickProductVersionForID(instanceId JoystickID) uint16 {
	return sdlGetJoystickProductVersionForID(instanceId)
}

func GetJoystickTypeForID(instanceId JoystickID) JoystickType {
	return sdlGetJoystickTypeForID(instanceId)
}

func OpenJoystick(instanceId JoystickID) *Joystick {
	return sdlOpenJoystick(instanceId)
}

func GetJoystickFromID(instanceId JoystickID) *Joystick {
	return sdlGetJoystickFromID(instanceId)
}

func GetJoystickFromPlayerIndex(playerIndex int32) *Joystick {
	return sdlGetJoystickFromPlayerIndex(playerIndex)
}

// type VirtualJoystickTouchpadDesc struct {
// 	NFingers uint16
// 	Padding1 uint16
// 	Padding2 uint16
// 	Padding3 uint16
// }

// type VirtualJoystickSensorDesc struct {
// 	Type SensorType
// 	Rate float32
// }
//
// type VirtualJoystickDesc struct{}

// func AttachVirtualJoystick(desc *VirtualJoystickDesc) JoystickID {
// 	return sdlAttachVirtualJoystick(desc)
// }

// func DetachVirtualJoystick(instance_id JoystickID) bool {
// 	return sdlDetachVirtualJoystick(instance_id)
// }

// func IsJoystickVirtual(instance_id JoystickID) bool {
// 	return sdlIsJoystickVirtual(instance_id)
// }

// func SetJoystickVirtualAxis(joystick *Joystick, axis int32, value int16) bool {
// 	return sdlSetJoystickVirtualAxis(joystick, axis, value)
// }

// func SetJoystickVirtualBall(joystick *Joystick, ball int32, xrel int16, yrel int16) bool {
// 	return sdlSetJoystickVirtualBall(joystick, ball, xrel, yrel)
// }

// func SetJoystickVirtualButton(joystick *Joystick, button int32, down bool) bool {
// 	return sdlSetJoystickVirtualButton(joystick, button, down)
// }

// func SetJoystickVirtualHat(joystick *Joystick, hat int32, value uint8) bool {
// 	return sdlSetJoystickVirtualHat(joystick, hat, value)
// }

// func SetJoystickVirtualTouchpad(joystick *Joystick, touchpad int32, finger int32, down bool, x float32, y float32, pressure float32) bool {
// 	return sdlSetJoystickVirtualTouchpad(joystick, touchpad, finger, down, x, y, pressure)
// }

// func SendJoystickVirtualSensorData(joystick *Joystick, sensorType SensorType, sensor_timestamp uint64, data *float32, num_values int32) bool {
// 	return sdlSendJoystickVirtualSensorData(joystick, sensorType, sensor_timestamp, data, num_values)
// }

func GetJoystickProperties(joystick *Joystick) PropertiesID {
	return sdlGetJoystickProperties(joystick)
}

const (
	PropJoystickCapMonoLEDBoolean       = "SDL.joystick.cap.mono_led"
	PropJoystickCapRGBLEDBoolean        = "SDL.joystick.cap.rgb_led"
	PropJoystickCapPlayerLEDBoolean     = "SDL.joystick.cap.player_led"
	PropJoystickCapRumbleBoolean        = "SDL.joystick.cap.rumble"
	PropJoystickCapTriggerRumbleBoolean = "SDL.joystick.cap.trigger_rumble"
)

func GetJoystickName(joystick *Joystick) string {
	return sdlGetJoystickName(joystick)
}

func GetJoystickPath(joystick *Joystick) string {
	return sdlGetJoystickPath(joystick)
}

func GetJoystickPlayerIndex(joystick *Joystick) int32 {
	return sdlGetJoystickPlayerIndex(joystick)
}

func SetJoystickPlayerIndex(joystick *Joystick, playerIndex int32) bool {
	return sdlSetJoystickPlayerIndex(joystick, playerIndex)
}

// func GetJoystickGUID(joystick *Joystick) GUID {
// 	return sdlGetJoystickGUID(joystick)
// }

func GetJoystickVendor(joystick *Joystick) uint16 {
	return sdlGetJoystickVendor(joystick)
}

func GetJoystickProduct(joystick *Joystick) uint16 {
	return sdlGetJoystickProduct(joystick)
}

func GetJoystickProductVersion(joystick *Joystick) uint16 {
	return sdlGetJoystickProductVersion(joystick)
}

func GetJoystickFirmwareVersion(joystick *Joystick) uint16 {
	return sdlGetJoystickFirmwareVersion(joystick)
}

func GetJoystickSerial(joystick *Joystick) string {
	return sdlGetJoystickSerial(joystick)
}

func GetJoystickType(joystick *Joystick) JoystickType {
	return sdlGetJoystickType(joystick)
}

// func GetJoystickGUIDInfo(guid GUID, vendor *uint16, product *uint16, version *uint16, crc16 *uint16) {
// 	sdlGetJoystickGUIDInfo(guid, vendor, product, version, crc16)
// }

func JoystickConnected(joystick *Joystick) bool {
	return sdlJoystickConnected(joystick)
}

func GetJoystickID(joystick *Joystick) JoystickID {
	return sdlGetJoystickID(joystick)
}

func GetNumJoystickAxes(joystick *Joystick) int32 {
	return sdlGetNumJoystickAxes(joystick)
}

func GetNumJoystickBalls(joystick *Joystick) int32 {
	return sdlGetNumJoystickBalls(joystick)
}

func GetNumJoystickHats(joystick *Joystick) int32 {
	return sdlGetNumJoystickHats(joystick)
}

func GetNumJoystickButtons(joystick *Joystick) int32 {
	return sdlGetNumJoystickButtons(joystick)
}

func SetJoystickEventsEnabled(enabled bool) {
	sdlSetJoystickEventsEnabled(enabled)
}

func JoystickEventsEnabled() bool {
	return sdlJoystickEventsEnabled()
}

func UpdateJoysticks() {
	sdlUpdateJoysticks()
}

func GetJoystickAxis(joystick *Joystick, axis int32) int16 {
	return sdlGetJoystickAxis(joystick, axis)
}

func GetJoystickAxisInitialState(joystick *Joystick, axis int32, state *int16) bool {
	return sdlGetJoystickAxisInitialState(joystick, axis, state)
}

func GetJoystickBall(joystick *Joystick, ball int32, dx *int32, dy *int32) bool {
	return sdlGetJoystickBall(joystick, ball, dx, dy)
}

func GetJoystickHat(joystick *Joystick, hat int32) uint8 {
	return sdlGetJoystickHat(joystick, hat)
}

const (
	HatCentered  = 0x00
	HatUp        = 0x01
	HatRight     = 0x02
	HatDown      = 0x04
	HatLeft      = 0x08
	HatRightUp   = HatRight | HatUp
	HatRightDown = HatRight | HatDown
	HatLeftUp    = HatLeft | HatUp
	HatLeftDown  = HatLeft | HatDown
)

func GetJoystickButton(joystick *Joystick, button int32) bool {
	return sdlGetJoystickButton(joystick, button)
}

func RumbleJoystick(joystick *Joystick, lowFrequencyRumble uint16, highFrequencyRumble uint16, durationMs uint32) bool {
	return sdlRumbleJoystick(joystick, lowFrequencyRumble, highFrequencyRumble, durationMs)
}

func RumbleJoystickTriggers(joystick *Joystick, leftRumble uint16, rightRumble uint16, durationMs uint32) bool {
	return sdlRumbleJoystickTriggers(joystick, leftRumble, rightRumble, durationMs)
}

func SetJoystickLED(joystick *Joystick, red uint8, green uint8, blue uint8) bool {
	return sdlSetJoystickLED(joystick, red, green, blue)
}

func SendJoystickEffect(joystick *Joystick, data unsafe.Pointer, size int32) bool {
	return sdlSendJoystickEffect(joystick, data, size)
}

func CloseJoystick(joystick *Joystick) {
	sdlCloseJoystick(joystick)
}

func GetJoystickConnectionState(joystick *Joystick) JoystickConnectionState {
	return sdlGetJoystickConnectionState(joystick)
}

func GetJoystickPowerInfo(joystick *Joystick, percent *int32) PowerState {
	return sdlGetJoystickPowerInfo(joystick, percent)
}
