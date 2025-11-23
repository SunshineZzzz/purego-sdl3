package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/SunshineZzzz/purego-sdl3/internal/convert"
)

type PropertyType uint32

const (
	PropertyTypeInvalid PropertyType = iota
	PropertyTypePointer
	PropertyTypeString
	PropertyTypeNumber
	PropertyTypeFloat
	PropertyTypeBoolean
)

type PropertiesID uint32

type CleanupPropertyCallback uintptr

func NewCleanupPropertyCallback(callback func(userdata, value unsafe.Pointer)) CleanupPropertyCallback {
	cb := purego.NewCallback(func(userdata, value unsafe.Pointer) uintptr {
		callback(userdata, value)
		return 0
	})

	return CleanupPropertyCallback(cb)
}

type EnumeratePropertiesCallback uintptr

func NewEnumeratePropertiesCallback(callback func(userdata unsafe.Pointer, props PropertiesID, name string)) EnumeratePropertiesCallback {
	cb := purego.NewCallback(func(userdata unsafe.Pointer, props PropertiesID, name *byte) uintptr {
		callback(userdata, props, convert.ToString(name))
		return 0
	})

	return EnumeratePropertiesCallback(cb)
}

func ClearProperty(props PropertiesID, name string) bool {
	return sdlClearProperty(props, name)
}

func CopyProperties(src PropertiesID, dst PropertiesID) bool {
	return sdlCopyProperties(src, dst)
}

func CreateProperties() PropertiesID {
	return sdlCreateProperties()
}

func DestroyProperties(props PropertiesID) {
	sdlDestroyProperties(props)
}

func EnumerateProperties(props PropertiesID, callback EnumeratePropertiesCallback, userdata unsafe.Pointer) bool {
	return sdlEnumerateProperties(props, callback, userdata)
}

func GetBooleanProperty(props PropertiesID, name string, defaultValue bool) bool {
	return sdlGetBooleanProperty(props, name, defaultValue)
}

func GetFloatProperty(props PropertiesID, name string, defaultValue float32) float32 {
	return sdlGetFloatProperty(props, name, defaultValue)
}

func GetGlobalProperties() PropertiesID {
	return sdlGetGlobalProperties()
}

func GetNumberProperty(props PropertiesID, name string, defaultValue int64) int64 {
	return sdlGetNumberProperty(props, name, defaultValue)
}

func GetPointerProperty(props PropertiesID, name string, defaultValue unsafe.Pointer) unsafe.Pointer {
	return sdlGetPointerProperty(props, name, defaultValue)
}

func GetPropertyType(props PropertiesID, name string) PropertyType {
	return sdlGetPropertyType(props, name)
}

func GetStringProperty(props PropertiesID, name string, defaultValue string) string {
	return sdlGetStringProperty(props, name, defaultValue)
}

func HasProperty(props PropertiesID, name string) bool {
	return sdlHasProperty(props, name)
}

func LockProperties(props PropertiesID) bool {
	return sdlLockProperties(props)
}

func SetBooleanProperty(props PropertiesID, name string, value bool) bool {
	return sdlSetBooleanProperty(props, name, value)
}

func SetFloatProperty(props PropertiesID, name string, value float32) bool {
	return sdlSetFloatProperty(props, name, value)
}

func SetNumberProperty(props PropertiesID, name string, value int64) bool {
	return sdlSetNumberProperty(props, name, value)
}

func SetPointerProperty(props PropertiesID, name string, value unsafe.Pointer) bool {
	return sdlSetPointerProperty(props, name, value)
}

func SetPointerPropertyWithCleanup(props PropertiesID, name string, value unsafe.Pointer, cleanup CleanupPropertyCallback, userdata unsafe.Pointer) bool {
	return sdlSetPointerPropertyWithCleanup(props, name, value, cleanup, userdata)
}

func SetStringProperty(props PropertiesID, name string, value string) bool {
	return sdlSetStringProperty(props, name, value)
}

func UnlockProperties(props PropertiesID) {
	sdlUnlockProperties(props)
}
